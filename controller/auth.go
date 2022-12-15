package controller

import (
	"net/http"

	"capstone-project/dto"
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func Connected(c echo.Context) error {
	return c.String(http.StatusOK, "Hello From Backend")
}

func Register(c echo.Context) error {
	var user entity.RegisterBinding
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	result, err := Service.Register(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	return c.JSON(http.StatusCreated, dto.BuildResponse(http.StatusCreated, http.StatusText(http.StatusCreated), result))
}

func Login(c echo.Context) error {
	var userBinding entity.LoginBinding
	if err := c.Bind(&userBinding); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	result, err := Service.Login(c, userBinding)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusCreated, http.StatusText(http.StatusCreated), result))
}

func GenerateRefreshToken(c echo.Context) error {
	var tokenBinding entity.TokenBinding
	if err := c.Bind(&tokenBinding); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	result, err := Service.ReGenerateToken(c, tokenBinding)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusCreated, http.StatusText(http.StatusCreated), result))
}