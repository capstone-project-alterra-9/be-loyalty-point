package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

<<<<<<< HEAD
func GetUsersPagination(c echo.Context) error {
	users, err := Service.GetUsersPagination(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", users))
}
=======
func GetOneByUserId(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.GetUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success to get user", user))
} 
>>>>>>> d1c1453a372f7fbfd69647ee36488e919f916ce1
