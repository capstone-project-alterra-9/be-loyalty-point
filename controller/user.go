package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteOneById(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.DeleteOneById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success to delete user", user))
} 