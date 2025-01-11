package data

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}
type User struct {
	ID             int64   `json:"id"`
	Username       string  `json:"username"`
	Birthday       string  `json:"birthday"`
	AwsCognitoId   string  `json:"aws_cognito_id"`
	CreatedAt      string  `json:"created_at"`
	Verified       bool    `json:"verified"`
	IsPrivate      bool    `json:"is_private"`
	InboxLocked    bool    `json:"inbox_locked"`
	SwiperMode     bool    `json:"swiper_mode"`
	Name           string  `json:"name,omitempty"`
	Gender         string  `json:"gender,omitempty"`
	CountryName    string  `json:"country_name,omitempty"`
	CountryFlag    string  `json:"country_flag,omitempty"`
	CountryIsoCode string  `json:"country_iso_code,omitempty"`
	CountryLat     float64 `json:"country_lat,omitempty"`
	CountryLng     float64 `json:"country_lng,omitempty"`
	CityName       string  `json:"city_name,omitempty"`
	CityLat        float64 `json:"city_lat,omitempty"`
	CityLng        float64 `json:"city_lng,omitempty"`
	Bio            string  `json:"bio,omitempty"`
	ProfileImageID int64   `json:"profile_image_id,omitempty"`
	ProfileImage   *Image  `json:"profile_image,omitempty"`
}

type UserPagination struct {
	Users      []*User `json:"users"`
	Total      int64   `json:"total"`
	HasMore    bool    `json:"has_more"`
	TotalPages int64   `json:"total_pages"`
	Page       int64   `json:"page"`
	Sort       string  `json:"sort"`
	Order      string  `json:"order"`
}

type UserRepositoryInterface interface {
	FindByUsername(username string) (*User, error)
	FindById(id int64) (*User, error)
	FindAll(page int64, limit int64, sort, order, search string) (*UserPagination, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*User, error) {
	query := `SELECT id, username, birthday, aws_cognito_id, created_at, verified, is_private,
			  inbox_locked, swiper_mode, COALESCE(name, ''), COALESCE(gender, ''),
			  COALESCE(country_name, ''), COALESCE(country_flag, ''), COALESCE(country_iso_code, ''),
			  COALESCE(country_lat, 0), COALESCE(country_lng, 0), COALESCE(city_name, ''),
			  COALESCE(city_lat, 0), COALESCE(city_lng, 0), COALESCE(bio, ''), COALESCE(profile_image_id, 0)
			  FROM users WHERE username = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt, &user.Verified,
		&user.IsPrivate, &user.InboxLocked, &user.SwiperMode, &user.Name, &user.Gender, &user.CountryName,
		&user.CountryFlag, &user.CountryIsoCode, &user.CountryLat, &user.CountryLng, &user.CityName,
		&user.CityLat, &user.CityLng, &user.Bio, &user.ProfileImageID)

	if err != nil {
		return nil, err
	}

	if user.ProfileImageID != 0 {
		query = `SELECT id, user_id, url, created_at FROM images WHERE id = $1`
		row = r.db.QueryRowContext(ctx, query, user.ProfileImageID)

		image := &Image{}
		err = row.Scan(&image.ID, &image.UserID, &image.URL, &image.CreatedAt)
		if err != nil {
			return nil, err
		}

		user.ProfileImage = image
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindById(id int64) (*User, error) {
	query := `SELECT id, username, birthday, aws_cognito_id, created_at, verified, is_private,
			  inbox_locked, swiper_mode, COALESCE(name, ''), COALESCE(gender, ''),
			  COALESCE(country_name, ''), COALESCE(country_flag, ''), COALESCE(country_iso_code, ''),
			  COALESCE(country_lat, 0), COALESCE(country_lng, 0), COALESCE(city_name, ''),
			  COALESCE(city_lat, 0), COALESCE(city_lng, 0), COALESCE(bio, ''), COALESCE(profile_image_id, 0)
			  FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt, &user.Verified,
		&user.IsPrivate, &user.InboxLocked, &user.SwiperMode, &user.Name, &user.Gender, &user.CountryName,
		&user.CountryFlag, &user.CountryIsoCode, &user.CountryLat, &user.CountryLng, &user.CityName,
		&user.CityLat, &user.CityLng, &user.Bio, &user.ProfileImageID)

	if err != nil {
		return nil, err
	}

	if user.ProfileImageID != 0 {
		query = `SELECT id, user_id, url, created_at FROM images WHERE id = $1`
		row = r.db.QueryRowContext(ctx, query, user.ProfileImageID)

		image := &Image{}
		err = row.Scan(&image.ID, &image.UserID, &image.URL, &image.CreatedAt)
		if err != nil {
			return nil, err
		}

		user.ProfileImage = image
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindAll(page int64, limit int64, sort, order, search string) (*UserPagination, error) {
	allowedSortFields := map[string]bool{
		"id":         true,
		"username":   true,
		"created_at": true,
	}
	if !allowedSortFields[sort] {
		sort = "id"
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	offset := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Base query parts
	whereClause := ""
	queryParams := make([]interface{}, 0)
	paramCount := 1

	// Add search condition if search parameter is provided
	if search != "" {
		whereClause = `
            WHERE (
                LOWER(username) LIKE $1 OR
                LOWER(COALESCE(name, '')) LIKE $1
            )`
		queryParams = append(queryParams, "%"+strings.ToLower(search)+"%")
		paramCount++
	}

	// Count total with search condition
	var total int64
	countQuery := `SELECT COUNT(*) FROM users ` + whereClause
	err := r.db.QueryRowContext(ctx, countQuery, queryParams...).Scan(&total)
	if err != nil {
		return nil, err
	}

	totalPages := (total + limit - 1) / limit

	// Main query with search condition
	query := `
        SELECT id, username, birthday, aws_cognito_id, created_at, verified, is_private,
               inbox_locked, swiper_mode, COALESCE(name, ''), COALESCE(gender, ''),
               COALESCE(country_name, ''), COALESCE(country_flag, ''), COALESCE(country_iso_code, ''),
               COALESCE(country_lat, 0), COALESCE(country_lng, 0), COALESCE(city_name, ''),
               COALESCE(city_lat, 0), COALESCE(city_lng, 0), COALESCE(bio, ''), COALESCE(profile_image_id, 0)
        FROM users
        ` + whereClause + `
        ORDER BY ` + sort + ` ` + order + `
        LIMIT $` + strconv.Itoa(paramCount) + ` OFFSET $` + strconv.Itoa(paramCount+1)

	// Add limit and offset to query params
	queryParams = append(queryParams, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		err := rows.Scan(
			&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt,
			&user.Verified, &user.IsPrivate, &user.InboxLocked, &user.SwiperMode,
			&user.Name, &user.Gender, &user.CountryName, &user.CountryFlag,
			&user.CountryIsoCode, &user.CountryLat, &user.CountryLng,
			&user.CityName, &user.CityLat, &user.CityLng, &user.Bio,
			&user.ProfileImageID,
		)
		if err != nil {
			return nil, err
		}

		if user.ProfileImageID != 0 {
			imageQuery := `SELECT id, user_id, url, created_at FROM images WHERE id = $1`
			row := r.db.QueryRowContext(ctx, imageQuery, user.ProfileImageID)

			image := &Image{}
			err = row.Scan(&image.ID, &image.UserID, &image.URL, &image.CreatedAt)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			if err != sql.ErrNoRows {
				user.ProfileImage = image
			}
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &UserPagination{
		Users:      users,
		Total:      total,
		HasMore:    page < totalPages,
		TotalPages: totalPages,
		Page:       page,
		Sort:       sort,
		Order:      order,
	}, nil
}
