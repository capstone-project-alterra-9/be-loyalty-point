package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOneByUserId(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.GetUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success to get user", user))
} 
