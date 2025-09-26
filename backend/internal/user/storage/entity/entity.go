package entity

import "gorm.io/gorm"

type User struct {
	Firstname string
	Lastname  string
	Age       int
	Email     string `gorm:"uniqueIndex;size:255"`
	gorm.Model
}
