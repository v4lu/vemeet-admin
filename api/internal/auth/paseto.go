package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/o1egl/paseto"
)

type TokenManager struct {
	paseto *paseto.V2
	secret string
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{
		paseto: paseto.NewV2(),
		secret: secret,
	}
}

func (tm *TokenManager) CreateToken(userID int64, expiration time.Time, tokenType string) (string, error) {
	claims := map[string]interface{}{
		"user_id": strconv.FormatInt(userID, 10),
		"exp":     expiration.Unix(),
		"type":    tokenType,
	}

	token, err := tm.paseto.Encrypt([]byte(tm.secret), claims, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return token, nil
}

func (tm *TokenManager) ValidateToken(token string, expectedType string) (int64, error) {
	var claims map[string]interface{}
	var footer string

	err := tm.paseto.Decrypt(token, []byte(tm.secret), &claims, &footer)
	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != expectedType {
		return 0, fmt.Errorf("invalid token type: expected '%s', got '%s'", expectedType, tokenType)
	}

	expClaim, ok := claims["exp"]
	if !ok {
		return 0, errors.New("missing expiration claim")
	}

	var expTime time.Time

	switch exp := expClaim.(type) {
	case float64:
		expTime = time.Unix(int64(exp), 0)
	case int64:
		expTime = time.Unix(exp, 0)
	default:
		return 0, errors.New("invalid expiration claim type")
	}

	if expTime.Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID format: %w", err)
	}

	return userID, nil
}
