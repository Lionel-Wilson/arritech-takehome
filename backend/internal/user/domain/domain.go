package domain

import "time"

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Age       int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
