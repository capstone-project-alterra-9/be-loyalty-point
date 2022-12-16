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
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
}

func GetUsersPagination(c echo.Context) error {
	users, err := Service.GetUsersPagination(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), users))
}

func GetOneByUserId(c echo.Context) error {
	id := c.Param("id")
	user, err := Service.GetUserById(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), user))
}

func UpdateOneByUserId(c echo.Context) error {
	id := c.Param("id")
	var user entity.UpdateUserBinding
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.UpdateOneById(c, id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func CreateUserByAdmin(c echo.Context) error {
	var user entity.CreateUserBinding
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	result, err := Service.CreateUserByAdmin(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetCountUsers(c echo.Context) error {
	result, err := Service.GetCountUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetForgotPassword(c echo.Context) error {
	var forgotPassword entity.ForgotPasswordBinding
	if err := c.Bind(&forgotPassword); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	err := Service.GetForgotPassword(c, forgotPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
}
