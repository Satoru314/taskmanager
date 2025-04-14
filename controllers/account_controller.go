package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccountsController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Accounts Handler\n")
}

func GetAccountController(c echo.Context) error {
	return c.String(http.StatusOK, "Get Account Handler\n")
}

func PostAccountController(c echo.Context) error {
	return c.String(http.StatusOK, "Post Account Handler\n")
}

func PutAccountController(c echo.Context) error {
	return c.String(http.StatusOK, "Put Account Handler\n")
}

func DeleteAccountController(c echo.Context) error {
	return c.String(http.StatusOK, "Delete Account Handler\n")
}
