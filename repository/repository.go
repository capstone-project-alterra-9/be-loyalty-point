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
	GetUserAuth(c echo.Context, user string) (*entity.Users, error)
	GetAdminAuth(c echo.Context, user string) (*entity.Admins, error)
}

type productRepo interface{}

type transactionRepo interface {
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsByUser(c echo.Context, ID string) ([]entity.Transactions, error)
}

type userRepo interface{}
