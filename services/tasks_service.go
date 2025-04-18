package services

import (
	"taskmanager/models"
	"taskmanager/repository"
)

func (s *MyAppServices) PostTaskService(reqTask models.Task) (resTask models.Task, err error) {
	resTask, err = repository.InsertTask(s.db, reqTask)
	return resTask, err
}

func (s *MyAppServices) GetTasksService(accountID int) (resTasks []models.Task, err error) {
	resTasks, err = repository.SelectTask(s.db, accountID)
	if err != nil {
		return []models.Task{}, err
	}
	return resTasks, nil
}

func (s *MyAppServices) PutTaskService(reqTask models.Task) (resTask models.Task, err error) {
	resTask, err = repository.UpdateTask(s.db, reqTask)
	return resTask, err
}

func (s *MyAppServices) DeleteTaskService(taskID int) (err error) {
	err = repository.DeleteTask(s.db, taskID)
	return err
}
