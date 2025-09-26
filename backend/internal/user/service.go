package user

import (
	"context"
	"fmt"

	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/mapper"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=user
type Service interface {
	CreateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, userID uint64) error
	GetUser(ctx context.Context, userID uint64) (domain.User, error)
}

type userService struct {
	logger   *logrus.Logger
	userRepo storage.UserRepository
}

func NewUserService(
	logger *logrus.Logger,
	userRepo storage.UserRepository,
) Service {
	return &userService{
		logger:   logger,
		userRepo: userRepo,
	}
}

var ErrUserMustBeAtLeast18YearsOld = fmt.Errorf("user must be at least 18 years old")

func (s *userService) CreateUser(ctx context.Context, user domain.User) error {
	if user.Age < 18 {
		return ErrUserMustBeAtLeast18YearsOld
	}

	err := s.userRepo.InsertUser(ctx, mapper.MapUserToEntity(user))
	if err != nil {
		return fmt.Errorf("failed to insert user into db: %w", err)
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, userID uint64) error {
	err := s.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user from db: %w", err)
	}

	return nil
}

func (s *userService) GetUser(ctx context.Context, userID uint64) (domain.User, error) {
	userEntity, err := s.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user from db: %w", err)
	}

	return mapper.MapUserEntityToDomain(userEntity), nil
}
