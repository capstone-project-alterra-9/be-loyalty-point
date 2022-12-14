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
	faqRepo
}

type authRepo interface {
	GetUserLogin(c echo.Context, user entity.LoginBinding) (*entity.Users, error)
	GetUserAuth(c echo.Context, user string) (*entity.Users, error)
	GetAdminAuth(c echo.Context, user string) (*entity.Users, error)
	GetAuth(c echo.Context, user string) (*entity.Users, error)
}

type productRepo interface {
	GetProductByID(c echo.Context, productID string) (*entity.Products, error)
	GetSerialNumbers(c echo.Context) ([]entity.SerialNumbers, error)
	GetSerialNumber(c echo.Context, productID string) (*entity.SerialNumbers, error)
	UpdateSerialStatus(c echo.Context, serial int64, status string) error
	CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error)
	CreateSerialNumber(c echo.Context, serialNumber *entity.SerialNumbers) error
	GetProducts(c echo.Context) ([]entity.Products, error)
	GetProductsByMethod(c echo.Context, paymentMethod string) ([]entity.Products, error)
	GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error)
	UpdateMethodProduct(c echo.Context, ID string, changeMethodBuy bool) error
	UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error)
	DeleteProduct(c echo.Context, ID string) error
	DeleteAllSerialNumberByProductID(c echo.Context, ID string) error
	DeleteNSerialNumberByProductID(c echo.Context, ID string, N int) error
	GetCountProducts(c echo.Context) (*entity.GetProductsCountView, error)
}

type transactionRepo interface {
	GetUserByIDRaw(c echo.Context, ID string) (entity.Users, error)
	GetProductByIDRaw(c echo.Context, ID string) (entity.Products, error)
	GetTransactions(c echo.Context) ([]entity.Transactions, error)
	GetTransactionsByMethod(c echo.Context, method string) ([]entity.Transactions, error)
	GetHistory(c echo.Context, ID string) ([]entity.Transactions, error)
	GetHistoryByMethod(c echo.Context, ID string, method string) ([]entity.Transactions, error)
	CreateTransaction(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error)
	GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error)
	UpdateTransaction(c echo.Context, ID string, transaction *entity.Transactions) (*entity.Transactions, error)
	DeleteTransaction(c echo.Context, ID string) error
	GetCountTransactions(c echo.Context) (*entity.GetTransactionsCountView, error)
}

type userRepo interface {
	CreateUser(c echo.Context, user entity.Users) (*entity.Users, error)
	CreatePoints(c echo.Context, userPoints entity.Points) (*entity.Points, error)
	GetUserPoints(c echo.Context, ID string) (*entity.Points, error)
	GetUserPointIgnoreEmpty(c echo.Context, ID string) (*entity.Points, error)
	UpdateUserPoints(c echo.Context, userPoint *entity.Points) error
	GetUsersPagination(c echo.Context) ([]entity.Users, error)
	GetUserByID(c echo.Context, ID string) (*entity.Users, error)
	UpdateOneByUserId(c echo.Context, user *entity.Users) (*entity.Users, error)
	DeleteUserById(c echo.Context, ID string) error
	GetCountUsers(c echo.Context) (*entity.GetUserCountView, error)
	DeleteUserPointsByUserId(c echo.Context, ID string) error
	GetUserByEmail(c echo.Context, email string) (*entity.Users, error)
}

type faqRepo interface {
	GetFaqs(c echo.Context) ([]entity.FAQ, error)
	CreateFaq(c echo.Context, user entity.FAQ) (*entity.FAQ, error)
}