package controller

import (
	"net/http"

	"capstone-project/dto"
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user entity.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	user.ID = (guuid.New()).String();
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	result, err := Service.Register(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to register user", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("succes register user", result))
}

func Login(c echo.Context) error {

	var user entity.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to process request", err))
	}

	result, err := Service.Login(c, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to login user", err))
	}

	return c.JSON(http.StatusOK, dto.BuildResponse("succes login user", result))
}
