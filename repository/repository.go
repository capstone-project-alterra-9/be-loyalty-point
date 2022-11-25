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
	GetUserAuth(c echo.Context) ([]entity.Users, error)
}

type productRepo interface {
	GetProduct(c echo.Context) ([]entity.Products, error)
	GetProductByID(c echo.Context, id uint) (entity.Products, error)
}

type transactionRepo interface {
	CreateTransaction(c echo.Context, transaction entity.Transactions) (entity.Transactions, error)
	GetSerialByStatus(c echo.Context, status string) (entity.SerialNumbers, error)
	UpdateSerialStatus(c echo.Context, updateResult entity.SerialNumbers) error
}

type userRepo interface{}
