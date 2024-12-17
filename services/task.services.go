package services

import (
	"lfernandez.com/todo/db"
	"lfernandez.com/todo/models"
)

var DB = db.DB

func GetTasks(limit int) []models.Task {
	var tasks []models.Task
	DB.Limit(limit).Find(&tasks)
	return tasks
}

func GetTaskById(id int) models.Task {
	var task models.Task
	DB.First(&task, id)
	return task
}

func CreateTask(tasks []models.Task) []models.Task {
	var createdTasks []models.Task
	for _, task := range tasks {
		DB.Create(&task)
		createdTasks = append(createdTasks, task)
	}
	return createdTasks
}

func UpdateTask(tasks []models.Task) []models.Task {
	var updatedTasks []models.Task
	for _, task := range tasks {
		DB.Save(&task)
		updatedTasks = append(updatedTasks, task)
	}
	return updatedTasks
}

func DeleteTask(tasks []models.Task) []models.Task {
	var deletedTasks []models.Task
	for _, task := range tasks {
		DB.Delete(&task)
		deletedTasks = append(deletedTasks, task)
	}
	return deletedTasks
}

func PatchTask(tasks []models.Task) []models.Task {
	var patchedTasks []models.Task
	for _, task := range tasks {
		DB.Save(&task)
		patchedTasks = append(patchedTasks, task)
	}
	return patchedTasks
}
