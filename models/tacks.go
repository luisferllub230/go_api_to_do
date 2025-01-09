package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      string
	Done        bool
	UserId      int
	User        User
}
