package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage/entity"
)

func MapUserToEntity(user domain.User) *entity.User {
	return &entity.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Email:     user.Email,
	}
}
