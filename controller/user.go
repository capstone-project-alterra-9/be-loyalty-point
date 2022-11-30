package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersPagination(c echo.Context) error {
	users, err := Service.GetUsersPagination(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", users))
}
