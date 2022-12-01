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
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type productController interface {
	CreateProduct(c echo.Context) error
	GetProducts(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
	GetProductByID(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type transactionController interface {
	GetTransactions(c echo.Context) error
	GetTransactionsByMethod(c echo.Context) error
	GetHistory(c echo.Context) error
	GetHistoryByMethod(c echo.Context) error
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
