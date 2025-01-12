package data

import (
	"context"
	"database/sql"
	"strconv"
	"time"
)

type Blocked struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at"`
	User      *User  `json:"user"`
}

type BlockedPagination struct {
	Blockeds   []*Blocked `json:"blockeds"`
	Total      int64      `json:"total"`
	HasMore    bool       `json:"has_more"`
	TotalPages int64      `json:"total_pages"`
	Page       int64      `json:"page"`
	Sort       string     `json:"sort"`
	Order      string     `json:"order"`
}

type BlockedRepositoryInterface interface {
	FindById(id int64) (*Blocked, error)
	FindAll(page int64, limit int64, sort, order, search string) (*BlockedPagination, error)
	Create(blocked *Blocked) (*Blocked, error)
	Update(blocked *Blocked) (*Blocked, error)
	Delete(id int64) (bool, error)
}

type BlockedRepositoryImpl struct {
	DB *sql.DB
}

func NewBlockedRepository(db *sql.DB) *BlockedRepositoryImpl {
	return &BlockedRepositoryImpl{DB: db}
}

func (r *BlockedRepositoryImpl) FindById(id int64) (*Blocked, error) {
	query := `SELECT id, user_id, reason, created_at FROM blockeds WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.DB.QueryRowContext(ctx, query, id)

	var blocked Blocked
	err := row.Scan(&blocked.ID, &blocked.UserID, &blocked.Reason, &blocked.CreatedAt)
	if err != nil {
		return nil, err
	}

	userQuery := userQuery()

	row = r.DB.QueryRowContext(ctx, userQuery, blocked.UserID)

	user := &User{}
	err = row.Scan(&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt, &user.Verified,
		&user.IsPrivate, &user.InboxLocked, &user.SwiperMode, &user.Blocked, &user.Name,
		&user.Gender, &user.CountryName, &user.CountryFlag, &user.CountryIsoCode, &user.CountryLat,
		&user.CountryLng, &user.CityName, &user.CityLat, &user.CityLng, &user.Bio, &user.ProfileImageID,
	)
	if err != nil {
		return nil, err
	}

	blocked.User = user

	return &blocked, nil
}

func (r *BlockedRepositoryImpl) Create(blocked *Blocked) (*Blocked, error) {
	query := `INSERT INTO blockeds (user_id, reason) VALUES ($1, $2) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.DB.QueryRowContext(ctx, query, blocked.UserID, blocked.Reason)

	err := row.Scan(&blocked.ID, &blocked.CreatedAt)
	if err != nil {
		return nil, err
	}

	userQuery := userQuery()

	row = r.DB.QueryRowContext(ctx, userQuery, blocked.UserID)

	user := &User{}
	err = row.Scan(&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt, &user.Verified,
		&user.IsPrivate, &user.InboxLocked, &user.SwiperMode, &user.Blocked, &user.Name,
		&user.Gender, &user.CountryName, &user.CountryFlag, &user.CountryIsoCode, &user.CountryLat,
		&user.CountryLng, &user.CityName, &user.CityLat, &user.CityLng, &user.Bio, &user.ProfileImageID,
	)
	if err != nil {
		return nil, err
	}

	blocked.User = user

	return blocked, nil
}

func (r *BlockedRepositoryImpl) Update(blocked *Blocked) (*Blocked, error) {
	query := `UPDATE blockeds SET user_id = $1, reason = $2 WHERE id = $3 RETURNING created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.DB.QueryRowContext(ctx, query, blocked.UserID, blocked.Reason, blocked.ID)

	err := row.Scan(&blocked.CreatedAt)
	if err != nil {
		return nil, err
	}

	userQuery := userQuery()

	row = r.DB.QueryRowContext(ctx, userQuery, blocked.UserID)

	user := &User{}
	err = row.Scan(&user.ID, &user.Username, &user.Birthday, &user.AwsCognitoId, &user.CreatedAt, &user.Verified,
		&user.IsPrivate, &user.InboxLocked, &user.SwiperMode, &user.Blocked, &user.Name,
		&user.Gender, &user.CountryName, &user.CountryFlag, &user.CountryIsoCode, &user.CountryLat,
		&user.CountryLng, &user.CityName, &user.CityLat, &user.CityLng, &user.Bio, &user.ProfileImageID,
	)
	if err != nil {
		return nil, err
	}

	blocked.User = user

	return blocked, nil
}

func (r *BlockedRepositoryImpl) Delete(id int64) (bool, error) {
	query := `DELETE FROM blockeds WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *BlockedRepositoryImpl) FindAll(page int64, limit int64, sort, order, search string) (*BlockedPagination, error) {
	allowedSortFields := map[string]bool{
		"id":         true,
		"user_id":    true,
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

	whereClause := ""
	queryParams := make([]interface{}, 0)
	paramCount := 1

	if search != "" {
		whereClause = `
						WHERE (
								LOWER(reason) LIKE $` + strconv.Itoa(paramCount) + ` OR
								LOWER(created_at) LIKE $` + strconv.Itoa(paramCount) + `
						)`
		queryParams = append(queryParams, "%"+search+"%")
		paramCount++
	}

	query := `SELECT id, user_id, reason, created_at FROM blockeds ` + whereClause + ` ORDER BY ` + sort + ` ` + order + ` LIMIT $` + strconv.Itoa(paramCount) + ` OFFSET $` + strconv.Itoa(paramCount+1)

	queryParams = append(queryParams, limit, offset)

	rows, err := r.DB.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blockeds := make([]*Blocked, 0)
	for rows.Next() {
		var blocked Blocked
		if err := rows.Scan(&blocked.ID, &blocked.UserID, &blocked.Reason, &blocked.CreatedAt); err != nil {
			return nil, err
		}
		blockeds = append(blockeds, &blocked)
	}

	countQuery := `SELECT COUNT(*) FROM blockeds ` + whereClause
	var total int64
	err = r.DB.QueryRowContext(ctx, countQuery, queryParams[:paramCount-1]...).Scan(&total)
	if err != nil {
		return nil, err
	}

	totalPages := (total + limit - 1) / limit
	hasMore := page < totalPages

	return &BlockedPagination{
		Blockeds:   blockeds,
		Total:      total,
		HasMore:    hasMore,
		TotalPages: totalPages,
		Page:       page,
		Sort:       sort,
		Order:      order,
	}, nil
}

func userQuery() string {
	return `SELECT id, username, birthday, aws_cognito_id, created_at, verified, is_private,
			  inbox_locked, swiper_mode, blocked, COALESCE(name, ''), COALESCE(gender, ''),
			  COALESCE(country_name, ''), COALESCE(country_flag, ''), COALESCE(country_iso_code, ''),
			  COALESCE(country_lat, 0), COALESCE(country_lng, 0), COALESCE(city_name, ''),
			  COALESCE(city_lat, 0), COALESCE(city_lng, 0), COALESCE(bio, ''), COALESCE(profile_image_id, 0)
			  FROM users WHERE username = $1`
}
