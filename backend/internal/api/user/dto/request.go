package dto

import "github.com/go-playground/validator"

type CreateUserRequest struct {
	Firstname   string `json:"firstname" validate:"required"`
	Lastname    string `json:"lastname" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
}

func (crr CreateUserRequest) Validate() error {
	return validator.New().Struct(crr)
}

type UpdateUserRequest struct {
	Firstname   *string `json:"firstname,omitempty"`
	Lastname    *string `json:"lastname,omitempty"`
	Age         *int    `json:"age,omitempty"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phonenumber,omitempty"`
}

func (uur UpdateUserRequest) Validate() error {
	return validator.New().Struct(uur)
}
