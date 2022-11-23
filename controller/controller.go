package controller

import (
	"capstone-project/service"
)

type Controller interface {
	authController
	productController
	transactionController
	userController
}

type authController interface{}

type productController interface{}

type transactionController interface{}

type userController interface{}

var Service service.Svc

func NewController(service service.Svc) {
	Service = service
}
