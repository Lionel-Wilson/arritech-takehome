package mapper

import "github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"

func MapUserToResponse(domain domain.User) User {
	return User{
		ID:        domain.ID,
		Firstname: domain.Firstname,
		Lastname:  domain.Lastname,
		Age:       domain.Age,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt.String(),
		UpdatedAt: domain.UpdatedAt.String(),
	}
}
