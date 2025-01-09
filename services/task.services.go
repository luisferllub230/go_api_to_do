package services

import (
	"lfernandez.com/todo/db"
	"lfernandez.com/todo/models"
)

func GetTasks(limit int) ([]models.Task, error) {
	var tasks []models.Task
	db.DB.Preload("User").Limit(limit).Find(&tasks)
	return tasks, nil
}

func GetTaskById(id int) models.Task {
	var task models.Task
	db.DB.Preload("User").First(&task, "id = ?", id)
	return task
}

func CreateTask(tasks []models.Task) ([]models.Task, error) {
	var createdTasks []models.Task
	for _, task := range tasks {
		createdTask := db.DB.Create(&task)
		createdTask.Preload("User")
		err := createdTask.Error

		if err != nil {
			return nil, err
		}

		createdTasks = append(createdTasks, task)
	}
	return createdTasks, nil
}

func UpdateTask(tasks []models.Task) ([]models.Task, error) {
	var updatedTasks []models.Task
	for _, task := range tasks {
		updatedTask := db.DB.Save(&task)
		err := updatedTask.Error

		if err != nil {
			return nil, err
		}
		updatedTasks = append(updatedTasks, task)

	}
	return updatedTasks, nil
}

func DeleteTask(tasks []models.Task) ([]models.Task, error) {
	var deletedTasks []models.Task
	for _, task := range tasks {
		deletedTask := db.DB.Delete(&task)
		err := deletedTask.Error

		if err != nil {
			return nil, err
		}

		deletedTasks = append(deletedTasks, task)
	}
	return deletedTasks, nil
}

func PatchTask(tasks []models.Task) []models.Task {
	var patchedTasks []models.Task
	for _, task := range tasks {
		db.DB.Save(&task)
		patchedTasks = append(patchedTasks, task)
	}
	return patchedTasks
}
