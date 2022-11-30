package service

import (
	"capstone-project/entity"
	"capstone-project/repository"

	"github.com/labstack/echo/v4"
)

type Service struct {
	repo repository.Repo
}

func NewService(repository repository.Repo) Svc {
	return &Service{
		repo: repository,
	}
}

type Svc interface {
	authSvc
	productSvc
	trasanctionSvc
	userSvc
}

type authSvc interface {
	Register(c echo.Context, user entity.RegisterBinding) (*entity.RegisterView, error)
	Login(c echo.Context, user entity.LoginBinding) (*entity.LoginView, error)
}

type productSvc interface{}

type trasanctionSvc interface {
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsRedeem(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsBuy(c echo.Context) ([]entity.Transactions, error)
	GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error)
	GetTransactionsByUser(c echo.Context) ([]entity.Transactions, error)
	CreateTransactionByUser(c echo.Context, transaction entity.TransactionsBinding) (*entity.Transactions, error)
	CreateTransactionByAdmin(c echo.Context, transaction entity.Transactions) (*entity.Transactions, error)
	UpdateTransactionByAdmin(c echo.Context, ID string, transaction entity.UpdateTransactionBinding) (*entity.Transactions, error)
	DeleteTransactionByAdmin(c echo.Context, transactionID string) error
}

type userSvc interface{}
