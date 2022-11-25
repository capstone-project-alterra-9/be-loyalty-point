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
	GetAdmins(c echo.Context) ([]entity.Admins, error)
}

type productRepo interface {
	CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error)
	CreateSerialNumber(c echo.Context, serialNumber *entity.SerialNumbers) error
	GetSerialNumbers(c echo.Context) ([]entity.SerialNumbers, error)
}

type transactionRepo interface{}

type userRepo interface{}
