package services

import (
	"taskmanager/models.go"
	"time"
)

func GetTaskService(accountID int) (resTask models.Task, err error) {
	// Simulate a service that retrieves a task by account ID
	// In a real application, this would involve database operations
	resTask = models.Task{
		TaskID:    1,
		AccountID: accountID,
		Title:     "Sample Task",
		DueDate:   time.Now(),
		Progress:  50,
	}
	return resTask, nil

}

func PostTaskService(reqTask models.Task) (resTask models.Task, err error) {
	// Simulate a service that creates a new task
	// In a real application, this would involve database operations
	resTask = reqTask
	resTask.TaskID = 1 // Simulate an auto-generated ID
	return resTask, nil
}

func PutTaskService(reqTask models.Task) (resTask models.Task, err error) {
	// Simulate a service that updates an existing task
	// In a real application, this would involve database operations
	resTask = reqTask
	return resTask, nil
}

func DeleteTaskService(taskID int) (err error) {
	// Simulate a service that deletes a task by ID
	// In a real application, this would involve database operations
	return nil
}

func GetTasksService(accountID int) (resTasks []models.Task, err error) {
	// Simulate a service that retrieves all tasks
	// In a real application, this would involve database operations
	resTasks = []models.Task{
		{
			TaskID:    1,
			AccountID: accountID,
			Title:     "Sample Task 1",
			DueDate:   time.Now(),
			Progress:  50,
		},
		{
			TaskID:    2,
			AccountID: accountID,
			Title:     "Sample Task 2",
			DueDate:   time.Now(),
			Progress:  75,
		},
	}
	return resTasks, nil
}
