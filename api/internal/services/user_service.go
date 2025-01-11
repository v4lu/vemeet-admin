package services

import (
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/errors"
)

type UserService struct {
	userRepo data.UserRepositoryInterface
}

type UserServiceInterface interface {
	GetUserByUsername(username string) (*data.User, error)
	GetUserById(id int64) (*data.User, error)
	GetUsers(page int64, limit int64, sort, order, search string) (*data.UserPagination, error)
}

func NewUserService(userRepo data.UserRepositoryInterface) UserServiceInterface {
	return &UserService{userRepo}
}

func (s *UserService) GetUserByUsername(username string) (*data.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.NewNotFoundError("user not found")
	}

	return user, nil
}

func (s *UserService) GetUserById(id int64) (*data.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.NewNotFoundError("user not found")
	}

	return user, nil
}

func (s *UserService) GetUsers(page int64, limit int64, sort, order, search string) (*data.UserPagination, error) {
	users, err := s.userRepo.FindAll(page, limit, sort, order, search)
	if err != nil {
		return nil, errors.NewInternalError("failed to get users")
	}

	return users, nil
}
