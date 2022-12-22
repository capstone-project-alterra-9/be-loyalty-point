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
	faqSvc
}

type authSvc interface {
	Register(c echo.Context, user entity.RegisterBinding) (*entity.RegisterView, error)
	Login(c echo.Context, user entity.LoginBinding) (*entity.LoginView, error)
	ReGenerateToken(c echo.Context, refreshToken entity.TokenBinding) (*entity.RefreshTokenView, error)
}

type productSvc interface {
	CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error)
	GetProducts(c echo.Context) ([]entity.Products, error)
	GetProductsByMethod(c echo.Context, paymentMethod string) ([]entity.Products, error)
	GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error)
	GetProductByID(c echo.Context, ID string) (*entity.Products, error)
	UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error)
	DeleteProduct(c echo.Context, ID string) error
	GetCountProducts(c echo.Context) (*entity.GetProductsCountView, error)
}

type trasanctionSvc interface {
	GetTransactions(c echo.Context) ([]entity.TransactionsView, error)
	GetTransactionsByMethod(c echo.Context, method string) ([]entity.TransactionsView, error)
	GetTransactionByID(c echo.Context, ID string) (*entity.TransactionsView, error)
	GetHistory(c echo.Context) ([]entity.TransactionsView, error)
	GetHistoryByMethod(c echo.Context, method string) ([]entity.TransactionsView, error)
	GetHistoryByMethodCategory(c echo.Context, method string, category string) ([]entity.TransactionsView, error)
	CreateTransactionByUser(c echo.Context, transaction entity.TransactionsBinding) (*entity.TransactionsView, error)
	CreateTransactionByAdmin(c echo.Context, transaction entity.Transactions) (*entity.TransactionsView, error)
	UpdateTransactionByAdmin(c echo.Context, ID string, transaction entity.UpdateTransactionBinding) (*entity.TransactionsView, error)
	DeleteTransactionByAdmin(c echo.Context, transactionID string) error
	GetCountTransactions(c echo.Context) (*entity.GetTransactionsCountView, error)
	CreateMidtransTransaction(c echo.Context, transaction entity.MidtransTransactionBinding) (*entity.MidtransTransactionView, error)
	HandlingPaymentStatus(c echo.Context, token string, status int) (string, error)
}

type userSvc interface {
	GetUsersPagination(c echo.Context, query entity.Paginate) (*entity.PaginationView, error)
	GetUserById(c echo.Context, ID string) (*entity.UsersView, error)
	UpdateOneById(c echo.Context, ID string, user entity.UpdateUserBinding) (*entity.UsersView, error)
	DeleteOneById(c echo.Context, userId string) error
	CreateUserByAdmin(c echo.Context, user entity.CreateUserBinding) (*entity.CreateUserView, error)
	GetCountUsers(c echo.Context) (*entity.GetUserCountView, error)
	SendEmailForgotPassword(c echo.Context, email entity.ForgotPasswordBinding) error
}

type faqSvc interface {
	GetFaqs(c echo.Context) ([]entity.FAQ, error)
	CreateFaq(c echo.Context, faq entity.FAQPayloadBinding) (*entity.FAQ, error)
}

