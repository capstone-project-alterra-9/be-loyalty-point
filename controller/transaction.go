package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTransaction(c echo.Context) error {
	var transactionParam entity.TransactionBinding
	err := c.Bind(&transactionParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Invalid request", err))
	}

	var result entity.Transactions
	result, err = Service.CreateTransaction(c, transactionParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to create transaction", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("Success create transaction", result))
}
