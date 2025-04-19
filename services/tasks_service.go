package services

import (
	"taskmanager/models"
	"taskmanager/repositorys"
)

func (s *MyAppServices) PostTaskService(reqTask models.Task) (resTask models.Task, err error) {
	resTask, err = repositorys.InsertTask(s.db, reqTask)
	return resTask, err
}

func (s *MyAppServices) GetTasksService(accountID int) (resTasks []models.Task, err error) {
	resTasks, err = repositorys.SelectTask(s.db, accountID)
	if err != nil {
		return []models.Task{}, err
	}
	return resTasks, nil
}

func (s *MyAppServices) PutTaskService(reqTask models.Task) (resTask models.Task, err error) {
	resTask, err = repositorys.UpdateTask(s.db, reqTask)
	return resTask, err
}

func (s *MyAppServices) DeleteTaskService(taskID int) (err error) {
	err = repositorys.DeleteTask(s.db, taskID)
	return err
}
