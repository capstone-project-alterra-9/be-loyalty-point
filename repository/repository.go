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
	GetUserByUsername(c echo.Context, username string) (*entity.Users, error)
}

type productRepo interface{}

type transactionRepo interface {
	CreateTransaction(c echo.Context, transaction entity.Transactions) (*entity.Transactions, error)
}

type userRepo interface{}
