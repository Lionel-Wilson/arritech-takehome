package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/api/user/dto"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
)

func MapUsersToResponse(domainUsers []domain.User, page, size int, total int64) dto.GetUsersResponse {
	var users []dto.User
	for _, user := range domainUsers {
		users = append(users, mapUserToDto(user))
	}
	return dto.GetUsersResponse{
		Users: users,
		Total: total,
		Page:  page,
		Size:  size,
	}
}

func MapUserToResponse(domain domain.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		User: mapUserToDto(domain)}
}

func mapUserToDto(user domain.User) dto.User {
	return dto.User{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Age:         user.Age,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.String(),
		UpdatedAt:   user.UpdatedAt.String(),
	}
}
