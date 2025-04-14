package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCommentsController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Comments Handler\n")
}

func GetCommentController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Comment Handler\n")
}

func PostCommentController(c echo.Context) error {
	return c.String(http.StatusOK, "Post Comment Handler\n")
}

func PutCommentController(c echo.Context) error {
	return c.String(http.StatusOK, "Put Comment Handler\n")
}

func DeleteCommentController(c echo.Context) error {
	return c.String(http.StatusOK, "Delete Comment Handler\n")
}
