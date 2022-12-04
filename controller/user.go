package controller

import (
	"capstone-project/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

<<<<<<< HEAD
func DeleteOneById(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.DeleteOneById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success to delete user", user))
} 
=======
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
>>>>>>> 344fb1ef0bd0d3eccd4bb333f6b6d2d64a329e93
