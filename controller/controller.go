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

type transactionController interface {
	GetTransactions(c echo.Context) error
	GetTransactionsRedeem(c echo.Context) error
	GetTransactionsBuy(c echo.Context) error
	GetTransactionsByUser(c echo.Context) error
	GetTransactionByID(c echo.Context) error
	CreateTransactionByUser(c echo.Context) error
	CreateTransactionByAdmin(c echo.Context) error
	UpdateTransactionByAdmin(c echo.Context) error
	DeleteTransactionByAdmin(c echo.Context) error
}

type userController interface{}

var Service service.Svc

func NewController(service service.Svc) {
	Service = service
}
