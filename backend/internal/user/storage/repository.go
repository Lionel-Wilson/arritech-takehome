package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=storage
type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

var ErrUserEmailExists = errors.New("user email already exists")

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
