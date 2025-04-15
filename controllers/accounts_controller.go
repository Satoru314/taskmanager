package controllers

import (
	"net/http"
	"strconv"

	"taskmanager/models.go"
	"taskmanager/services.go"

	"github.com/labstack/echo/v4"
)

func GetAccountHandler(c echo.Context) error {
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account ID\n")
	}
	resAccounts, err := services.GetAccountService(accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error retrieving accounts\n")
	}

	return c.JSON(http.StatusOK, resAccounts)
}

func PostAccountHandler(c echo.Context) error {
	var reqAccount models.Account
	c.Bind(&reqAccount)
	resAccount, err := services.PostAccountService(reqAccount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating account\n")
	}

	return c.JSON(http.StatusOK, resAccount)
}

func PutAccountHandler(c echo.Context) error {
	var reqAccount models.Account
	if err := c.Bind(&reqAccount); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload\n")
	}
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account ID\n")
	}
	reqAccount.AccountID = accountID
	resAccount, err := services.PutAccountService(reqAccount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating account\n")
	}

	return c.JSON(http.StatusOK, resAccount)
}

func DeleteAccountHandler(c echo.Context) error {
	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid account ID\n")
	}
	err = services.DeleteAccountService(accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting account\n")
	}

	return c.JSON(http.StatusOK, "Delete Account Successed!")
}
