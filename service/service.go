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

type authSvc interface{}

type productSvc interface{}

type trasanctionSvc interface {
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsByCategories(c echo.Context, method string) ([]entity.Transactions, error)
	GetTransactionsByUser(c echo.Context) ([]entity.Transactions, error)
}

type userSvc interface{}
