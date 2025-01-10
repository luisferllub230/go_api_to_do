package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"required;not null"`
	Task     []Task `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
