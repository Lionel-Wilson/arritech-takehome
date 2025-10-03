package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/api/user/dto"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
)

func MapCreateUserRequestToDomain(req dto.CreateUserRequest) domain.User {
	return domain.User{
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
}

func MapUpdateUserRequestToDomain(req dto.UpdateUserRequest, userID uint64) domain.UpdateUser {
	return domain.UpdateUser{
		ID:          userID,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Age:         req.Age,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
}
