package controller

import (
	"net/http"

	"capstone-project/dto"
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user entity.RegisterBinding
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	result, err := Service.Register(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to register user", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("succes register user", result))
}

func Login(c echo.Context) error {
	var userBinding entity.LoginBinding
	if err := c.Bind(&userBinding); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	result, err := Service.Login(c, userBinding)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to login user", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("succes login user", result))
}
