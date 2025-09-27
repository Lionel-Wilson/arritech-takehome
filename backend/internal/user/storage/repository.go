package storage

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/Lionel-Wilson/arritech-takehome/internal/entity"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=storage
type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID uint64) error
	GetUserById(ctx context.Context, userID uint64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context, p GetUsersParams) ([]entity.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

var (
	ErrUserEmailExists = errors.New("user email already exists")
	ErrUserNotFound    = errors.New("user not found")
)

type GetUsersParams struct {
	Query    string
	Page     int
	PageSize int
}

func (r *userRepository) GetUsers(ctx context.Context, p GetUsersParams) ([]entity.User, int64, error) {
	db := r.db.WithContext(ctx).Model(&entity.User{}).Where("deleted_at IS NULL")

	if q := strings.TrimSpace(p.Query); q != "" {
		like := "%" + strings.ToLower(q) + "%"

		// Base text search
		cond := r.db.Where(
			"LOWER(firstname) LIKE ? OR LOWER(lastname) LIKE ? OR LOWER(email) LIKE ?",
			like, like, like,
		)

		// If the query is numeric, allow searching by exact user id too.
		if id, err := strconv.Atoi(q); err == nil {
			cond = cond.Or("id = ?", id)
		}

		db = db.Where(cond)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}
	offset := (p.Page - 1) * p.PageSize

	var users []entity.User
	err := db.
		Order("id ASC").
		Limit(p.PageSize).
		Offset(offset).
		Find(&users).Error

	return users, total, err
}
func (r *userRepository) InsertUser(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // unique_violation
			return ErrUserEmailExists
		}

		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userID uint64) error {
	err := r.db.WithContext(ctx).Unscoped().Delete(&entity.User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserById(ctx context.Context, userID uint64) (*entity.User, error) {
	var user *entity.User

	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return err
	}

	return nil
}
