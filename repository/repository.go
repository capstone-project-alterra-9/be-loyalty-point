package repository

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repository struct {
	connection *gorm.DB
}

func NewRepository(db *gorm.DB) Repo {
	return &repository{
		connection: db,
	}
}

type Repo interface {
	authRepo
	productRepo
	transactionRepo
	userRepo
}

type authRepo interface {
	Register(c echo.Context, user entity.Users) (*entity.Users, error)
}

type productRepo interface{}

type transactionRepo interface{}

type userRepo interface{}
