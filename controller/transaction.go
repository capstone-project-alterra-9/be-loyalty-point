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

func GetTransactionsByCategories(c echo.Context) error {
	method := c.QueryParam("method")
	transactions, err := Service.GetTransactionsByCategories(c, method)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", transactions))
}

func GetTransactionsByUser(c echo.Context) error {
	transactions, err := Service.GetTransactionsByUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", transactions))
}
