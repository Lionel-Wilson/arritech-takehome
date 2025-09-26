package mapper

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/entity"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
)

func MapUserToEntity(user domain.User) *entity.User {
	userEntity := &entity.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Email:     user.Email,
	}

	if user.ID != 0 {
		userEntity.ID = user.ID
	}

	if !user.CreatedAt.IsZero() {
		userEntity.CreatedAt = user.CreatedAt
	}
	return userEntity
}
