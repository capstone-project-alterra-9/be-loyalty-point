package controller

import (
	"capstone-project/dto"
	"capstone-project/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	// Get user id from token
	user := middleware.ExtractTokenUsername(c)
	transactions, err := Service.GetTransactions(c, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", transactions))
}
