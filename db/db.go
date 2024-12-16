package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func InitConnection(config Config) *gorm.DB {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.Host, config.User, config.Password, config.Name, config.Port)
	db, error := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if error != nil {
		panic(error)
	} else {
		fmt.Println("Connected to database")
	}

	return db
}
