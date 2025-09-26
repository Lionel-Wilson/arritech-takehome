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
	UpdateUserDetails(ctx context.Context, updatedUser domain.UpdateUser) error
	GetUsers(ctx context.Context) ([]domain.User, error)
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

func (s *userService) GetUsers(ctx context.Context) ([]domain.User, error) {
	userEntities, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users from db: %w", err)
	}

	var domainUsers []domain.User
	for _, userEntity := range userEntities {
		domainUsers = append(domainUsers, mapper.MapUserEntityToDomain(&userEntity))
	}

	return domainUsers, nil
}

func (s *userService) UpdateUserDetails(ctx context.Context, updatedUser domain.UpdateUser) error {
	userDetails, err := s.GetUser(ctx, updatedUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get user details: %w", err)
	}

	if updatedUser.Firstname != nil {
		userDetails.Firstname = *updatedUser.Firstname
	}
	if updatedUser.Lastname != nil {
		userDetails.Lastname = *updatedUser.Lastname
	}
	if updatedUser.Age != nil {
		if *updatedUser.Age < 18 {
			return ErrUserMustBeAtLeast18YearsOld
		}
		userDetails.Age = *updatedUser.Age
	}
	if updatedUser.Email != nil {
		userDetails.Email = *updatedUser.Email
	}

	err = s.userRepo.UpdateUser(ctx, mapper.MapUserToEntity(userDetails))
	if err != nil {
		return fmt.Errorf("failed to update user details: %w", err)
	}
	return nil

}

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
