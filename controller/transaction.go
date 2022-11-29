package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	transactions, err := Service.GetTransactions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", transactions))
}

func GetTransactionsRedeem(c echo.Context) error {
	transactions, err := Service.GetTransactionsRedeem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions redeem", transactions))
}

func GetTransactionsBuy(c echo.Context) error {
	transactions, err := Service.GetTransactionsBuy(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions buy", transactions))
}

func GetTransactionsByUser(c echo.Context) error {
	transactions, err := Service.GetTransactionsByUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions user", transactions))
}
