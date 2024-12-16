package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lfernandez.com/todo/models"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	DBType   string
}

var DB *gorm.DB

func InitConnection(config Config) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.Host, config.User, config.Password, config.Name, config.Port)
	DB, error := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if error != nil {
		panic(error)
	} else {
		fmt.Println("Connected to database")
	}

	DB.AutoMigrate(models.User{}, models.Task{})
}
