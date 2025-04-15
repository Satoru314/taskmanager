package controllers

import (
	"net/http"
	"strconv"

	"taskmanager/models"
	"taskmanager/services"

	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account ID\n")
	}
	resTasks, err := services.GetTasksService(accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error retrieving tasks\n")
	}

	return c.JSON(http.StatusOK, resTasks)
}

func PostTaskHandler(c echo.Context) error {
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account ID\n")
	}
	var reqTask models.Task
	c.Bind(&reqTask)
	reqTask.AccountID = accountID
	resTask, err := services.PostTaskService(reqTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating task\n")
	}

	return c.JSON(http.StatusOK, resTask)
}

func PutTaskHandler(c echo.Context) error {
	taskIDStr := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task ID\n")
	}
	var reqTask models.Task
	c.Bind(&reqTask)
	reqTask.TaskID = taskID
	resTask, err := services.PutTaskService(reqTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating task\n")
	}

	return c.JSON(http.StatusOK, resTask)
}

func DeleteTaskHandler(c echo.Context) error {
	taskIDStr := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid task ID\n")
	}
	err = services.DeleteTaskService(taskID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting task\n")
	}

	return c.JSON(http.StatusOK, "Delete Task Succeeded!")
}
