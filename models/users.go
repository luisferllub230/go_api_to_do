package models

import "gorm.io/gorm"

// TODO: asociate with tasks

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
