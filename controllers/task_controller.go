package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasksController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Tasks Handler\n")
}

func GetTaskController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Task Handler\n")
}

func PostTaskController(c echo.Context) error {
	return c.String(http.StatusOK, "Post Task Handler\n")
}

func PutTaskController(c echo.Context) error {
	return c.String(http.StatusOK, "Put Task Handler\n")
}

func DeleteTaskController(c echo.Context) error {
	return c.String(http.StatusOK, "Delete Task Handler\n")
}
