package models

import "gorm.io/gorm"

// TODO: asociate with tasks

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Task     []Task `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
