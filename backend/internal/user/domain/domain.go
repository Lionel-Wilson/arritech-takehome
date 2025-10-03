package domain

import "time"

type User struct {
	ID          uint
	Firstname   string
	Lastname    string
	Age         int
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateUser struct {
	ID          uint64
	Firstname   *string
	Lastname    *string
	Age         *int
	Email       *string
	PhoneNumber *string
}

type GetUsersParams struct {
	Query string
	Page  int
	Size  int
}
