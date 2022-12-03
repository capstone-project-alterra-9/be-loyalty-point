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

type productSvc interface {
	CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error)
	GetProducts(c echo.Context) ([]entity.Products, error)
	GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error)
	GetProductByID(c echo.Context, ID string) (*entity.Products, error)
	UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error)
	DeleteProduct(c echo.Context, ID string) error
}

type trasanctionSvc interface {
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsByMethod(c echo.Context, method string) ([]entity.Transactions, error)
	GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error)
	GetHistory(c echo.Context) ([]entity.Transactions, error)
	GetHistoryByMethod(c echo.Context, method string) ([]entity.Transactions, error)
	CreateTransactionByUser(c echo.Context, transaction entity.TransactionsBinding) (*entity.Transactions, error)
	CreateTransactionByAdmin(c echo.Context, transaction entity.Transactions) (*entity.Transactions, error)
	UpdateTransactionByAdmin(c echo.Context, ID string, transaction entity.UpdateTransactionBinding) (*entity.Transactions, error)
	DeleteTransactionByAdmin(c echo.Context, transactionID string) error
}

type userSvc interface{
	GetUserById(c echo.Context, ID string) (*entity.Users, error)
}
