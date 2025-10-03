package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/entity"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
)

func MapUserEntityToDomain(ent *entity.User) domain.User {
	return domain.User{
		ID:          ent.ID,
		Firstname:   ent.Firstname,
		Lastname:    ent.Lastname,
		Age:         ent.Age,
		Email:       ent.Email,
		PhoneNumber: ent.PhoneNumber,
		CreatedAt:   ent.CreatedAt,
		UpdatedAt:   ent.UpdatedAt,
	}
}
