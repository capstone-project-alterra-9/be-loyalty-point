package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteOneById(c echo.Context) error {
	id := c.Param("id")
	err := Service.DeleteOneById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success delete user", dto.EmptyObj{}))
}

func GetUsersPagination(c echo.Context) error {
	users, err := Service.GetUsersPagination(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get all user", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get all user", users))
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
	var user entity.UpdateUserBinding
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

func CreateUserByAdmin(c echo.Context) error {
	var user entity.CreateUserBinding
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	result, err := Service.CreateUserByAdmin(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to create user", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("succes register user", result))
}

func GetCountUsers(c echo.Context) error {
	result, err := Service.GetCountUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get total users", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get total users", result))
}

func GetForgotPassword(c echo.Context) error {
	var forgotPassword entity.ForgotPasswordBinding
	if err := c.Bind(&forgotPassword); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	err := Service.GetForgotPassword(c, forgotPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get total users", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get total users", ""))
}
