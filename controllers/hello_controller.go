package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!\n")
}
