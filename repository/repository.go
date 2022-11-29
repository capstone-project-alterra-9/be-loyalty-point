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
	GetUserLogin(c echo.Context, user entity.LoginBinding) (*entity.Users, error)
	GetUserAuth(c echo.Context, user string) (*entity.Users, error)
	GetAdminAuth(c echo.Context, user string) (*entity.Users, error)
}

type productRepo interface{}

type transactionRepo interface {
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsRedeem(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsBuy(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsByUser(c echo.Context, ID string) ([]entity.Transactions, error)
}

type userRepo interface {
	CreateUser(c echo.Context, user entity.Users) (*entity.Users, error)
	CreatePoints(c echo.Context, userPoints entity.Points) (*entity.Points, error)
}
