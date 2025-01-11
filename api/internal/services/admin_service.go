package services

import (
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/errors"
)

type AdminService struct {
	adminRepo data.AdminRepositoryInterface
}

type AdminServiceInterface interface {
	FindAdminByEmail(email string) (*data.Admin, error)
	FindAdminById(id int64) (*data.Admin, error)
	InsertAdmin(admin *data.Admin) error
	UpdateAdmin(admin *data.Admin) error
	FindAllAdmins() ([]*data.Admin, error)
}

func NewAdminService(adminRepo data.AdminRepositoryInterface) AdminServiceInterface {
	return &AdminService{adminRepo}
}

func (s *AdminService) FindAdminByEmail(email string) (*data.Admin, error) {
	if email == "" {
		return nil, errors.NewValidationError("email name is required")
	}

	admin, err := s.adminRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.NewNotFoundError("admin not found")
	}

	return admin, nil
}

func (s *AdminService) FindAdminById(id int64) (*data.Admin, error) {
	if id == 0 {
		return nil, errors.NewValidationError("id is required")
	}

	admin, err := s.adminRepo.FindById(id)
	if err != nil {
		return nil, errors.NewNotFoundError("admin not found")
	}

	return admin, nil
}

func (s *AdminService) InsertAdmin(admin *data.Admin) error {
	if admin.Email == "" || admin.Password == "" || admin.Name == "" {
		return errors.NewValidationError("email is required")
	}

	_, err := s.FindAdminByEmail(admin.Email)
	if err == nil {
		return errors.NewValidationError("email already exists")
	}

	hashedPassword, err := HashPassword(admin.Password)
	if err != nil {
		return errors.NewValidationError("password not hashed")
	}

	admin.Password = hashedPassword

	err = s.adminRepo.InserAdmin(admin)
	if err != nil {
		return errors.NewValidationError("admin not inserted")
	}

	return nil
}

func (s *AdminService) UpdateAdmin(admin *data.Admin) error {
	if admin == nil {
		return errors.NewValidationError("admin is required")
	}

	err := s.adminRepo.UpdateAdmin(admin)
	if err != nil {
		return errors.NewValidationError("admin not updated")
	}

	return nil
}

func (s *AdminService) FindAllAdmins() ([]*data.Admin, error) {
	admins, err := s.adminRepo.FindAll()
	if err != nil {
		return nil, errors.NewNotFoundError("admins not found")
	}

	return admins, nil
}
