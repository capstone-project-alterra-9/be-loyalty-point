package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTransaction(c echo.Context) error {
	var transaction entity.Transactions
	err := c.Bind(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Invalid request", err))
	}

	var result entity.Transactions
	result, err = Service.CreateTransaction(c, transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to create transaction", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("Success create transaction", result))
}
