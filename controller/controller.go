package controller

import (
	"capstone-project/service"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	authController
	productController
	transactionController
	userController
}

type authController interface {
	Connected(c echo.Context) error
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type productController interface{}

type transactionController interface{}

type userController interface{}

var Service service.Svc

func NewController(service service.Svc) {
	Service = service
}
