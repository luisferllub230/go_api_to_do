package services

import (
	"lfernandez.com/todo/db"
	"lfernandez.com/todo/models"
)

func GetUsers(limit int) ([]models.User, error) {
	var users []models.User
	db.DB.Preload("Task").Limit(limit).Find(&users)
	return users, nil
}

func CreateUsers(users []models.User) ([]models.User, error) {
	var createdUsers []models.User
	for _, user := range users {
		createdUser := db.DB.Create(&user)
		createdUser.Preload("Task")
		err := createdUser.Error

		if err != nil {
			return nil, err
		}

		createdUsers = append(createdUsers, user)
	}
	return createdUsers, nil
}

func UpdateUsers(users []models.User) ([]models.User, error) {
	var updatedUsers []models.User
	for _, user := range users {
		updatedUser := db.DB.Save(&user)
		err := updatedUser.Error

		if err != nil {
			return nil, err
		}
		updatedUsers = append(updatedUsers, user)

	}
	return updatedUsers, nil
}

func DeleteUsers(users []models.User) ([]models.User, error) {
	var deletedUsers []models.User
	for _, user := range users {
		deletedUser := db.DB.Delete(&user)
		err := deletedUser.Error

		if err != nil {
			return nil, err
		}

		deletedUsers = append(deletedUsers, user)
	}
	return deletedUsers, nil
}
