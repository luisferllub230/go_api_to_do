package models

import "gorm.io/gorm"

// TODO: asociate task with user

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      string
	Done        bool
	UserId      uint
}
