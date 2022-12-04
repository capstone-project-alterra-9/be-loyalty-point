package controller

import (
	"capstone-project/dto"
	"net/http"
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func GetUsersPagination(c echo.Context) error {
	users, err := Service.GetUsersPagination(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", users))
}

func GetOneByUserId(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.GetUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success to get user", user))
} 

func UpdateOneByUserId(c echo.Context) error {
	id := c.Param("id")
	var user entity.Users
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.UpdateOneById(c, id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to update user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success update user", result))
}