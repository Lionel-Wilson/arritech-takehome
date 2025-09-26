package storage

import (
	"context"
	"errors"
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
	GetUsers(ctx context.Context) ([]entity.User, error)
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

func (r *userRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
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
