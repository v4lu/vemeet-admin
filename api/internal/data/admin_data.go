package data

import (
	"context"
	"database/sql"
	"time"
)

type Admin struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Verified  bool   `json:"verified"`
	CreatedAt string `json:"created_at"`
}

type AdminRepositoryInterface interface {
	FindByEmail(email string) (*Admin, error)
	FindById(id int64) (*Admin, error)
	InserAdmin(admin *Admin) error
	UpdateAdmin(admin *Admin) error
	FindAll() ([]*Admin, error)
}

type AdminRepositoryImpl struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepositoryInterface {
	return &AdminRepositoryImpl{db}
}

func (r *AdminRepositoryImpl) FindByEmail(email string) (*Admin, error) {
	query := `SELECT id, email, password, name, verified, created_at FROM admin_users WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, email)

	var admin Admin
	err := row.Scan(&admin.ID, &admin.Email, &admin.Password, &admin.Name, &admin.Verified, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) FindById(id int64) (*Admin, error) {
	query := `SELECT id, email, password, name, verified, created_at FROM admin_users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, id)

	var admin Admin
	err := row.Scan(&admin.ID, &admin.Email, &admin.Password, &admin.Name, &admin.Verified, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) InserAdmin(admin *Admin) error {
	query := `INSERT INTO admin_users (email, password, name) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, admin.Email, admin.Password, admin.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepositoryImpl) UpdateAdmin(admin *Admin) error {
	query := `UPDATE admin_users SET email = $1, password = $2, name = $3, verified = $4 WHERE id = $5`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, admin.Email, admin.Password, admin.Name, admin.Verified, admin.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepositoryImpl) FindAll() ([]*Admin, error) {
	query := `SELECT id, email, password, name, verified, created_at FROM admin_users`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []*Admin
	for rows.Next() {
		var admin Admin
		err := rows.Scan(&admin.ID, &admin.Email, &admin.Password, &admin.Name, &admin.Verified, &admin.CreatedAt)
		if err != nil {
			return nil, err
		}

		admins = append(admins, &admin)
	}

	return admins, nil
}
