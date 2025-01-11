package services

import (
	"context"
	"fmt"
	"time"

	"github.com/valu/vemeet-admin-api/internal/auth"
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/errors"
	"github.com/valu/vemeet-admin-api/internal/models"
)

type AuthService struct {
	adminRepo    data.AdminRepositoryInterface
	tokenManager auth.TokenManager
}

type AuthServiceInterface interface {
	LoginUser(email, password string) (*models.TokenPair, *data.Admin, error)
	RefreshTokens(refreshToken string) (*models.TokenPair, error)
	GetSession(userId int64) (*data.Admin, error)
}

func NewAuthService(adminRepo data.AdminRepositoryInterface, tokenManager auth.TokenManager) AuthServiceInterface {
	return &AuthService{adminRepo, tokenManager}
}

func (s *AuthService) LoginUser(email, password string) (*models.TokenPair, *data.Admin, error) {
	if email == "" || password == "" {
		return nil, nil, errors.NewValidationError("email and password are required")
	}

	admin, err := s.adminRepo.FindByEmail(email)
	if err != nil {
		return nil, nil, errors.NewNotFoundError("admin not found")
	}

	if !admin.Verified {
		return nil, nil, errors.NewValidationError("admin is not verified")
	}

	err = ComparePassword(admin.Password, password)
	if err != nil {
		return nil, nil, errors.NewValidationError("password is incorrect")
	}

	tokens, err := s.GenerateTokens(context.Background(), admin.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return tokens, admin, nil
}

func (s *AuthService) RefreshTokens(refreshToken string) (*models.TokenPair, error) {
	userID, err := s.tokenManager.ValidateToken(refreshToken, "refresh")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	tokens, err := s.GenerateTokens(context.Background(), userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return tokens, nil
}

func (s *AuthService) GenerateTokens(ctx context.Context, userID int64) (*models.TokenPair, error) {
	now := time.Now()
	accessTokenExp := now.Add(60 * time.Minute)
	refreshTokenExp := now.Add(7 * 24 * time.Hour)

	accessToken, err := s.tokenManager.CreateToken(userID, accessTokenExp, "access")
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	refreshToken, err := s.tokenManager.CreateToken(userID, refreshTokenExp, "refresh")
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	return &models.TokenPair{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  accessTokenExp,
		RefreshTokenExpiry: refreshTokenExp,
	}, nil
}

func (s *AuthService) GetSession(userId int64) (*data.Admin, error) {

	admin, err := s.adminRepo.FindById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin: %w", err)
	}

	return admin, nil
}
