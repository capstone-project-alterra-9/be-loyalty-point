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

type authController interface{}

type productController interface{}

type transactionController interface {
	GetTransactions(c echo.Context) error
}

type userController interface{}

var Service service.Svc

func NewController(service service.Svc) {
	Service = service
}
