package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage/entity"
)

func MapUserEntityToDomain(ent *entity.User) domain.User {
	return domain.User{
		ID:        ent.ID,
		Firstname: ent.Firstname,
		Lastname:  ent.Lastname,
		Age:       ent.Age,
		Email:     ent.Email,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}
