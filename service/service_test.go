package service

// fakhri try manual mock to service unit test

import (
	"capstone-project/entity"
	// "capstone-project/repository"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	// allService                       Svc
	// allRepo                          repository.Repo
	getUserLogin                     func(c echo.Context, user entity.LoginBinding) (*entity.Users, error)
	getUserAuth                      func(c echo.Context, user string) (*entity.Users, error)
	getAdminAuth                     func(c echo.Context, user string) (*entity.Users, error)
	getAuth                          func(c echo.Context, user string) (*entity.Users, error)
	getProductByID                   func(c echo.Context, productID string) (*entity.Products, error)
	getSerialNumbers                 func(c echo.Context) ([]entity.SerialNumbers, error)
	getSerialNumber                  func(c echo.Context, productID string) (*entity.SerialNumbers, error)
	updateSerialStatus               func(c echo.Context, serial int64, status string) error
	createProduct                    func(c echo.Context, product *entity.Products) (*entity.Products, error)
	createSerialNumber               func(c echo.Context, serialNumber *entity.SerialNumbers) error
	getProducts                      func(c echo.Context) ([]entity.Products, error)
	getProductsByMethod              func(c echo.Context, paymentMethod string) ([]entity.Products, error)
	getProductsByCategory            func(c echo.Context, category string) ([]entity.Products, error)
	updateCategoryProduct            func(c echo.Context, ID string, category string) error
	updateMethodProduct              func(c echo.Context, ID string, changeMethodBuy bool) error
	updateProduct                    func(c echo.Context, ID string, product *entity.Products) (*entity.Products, error)
	deleteProduct                    func(c echo.Context, ID string) error
	deleteAllSerialNumberByProductID func(c echo.Context, ID string) error
	deleteNSerialNumberByProductID   func(c echo.Context, ID string, N int) error
	getUserByIDRaw                   func(c echo.Context, ID string) (entity.Users, error)
	getProductByIDRaw                func(c echo.Context, ID string) (entity.Products, error)
	getTransactions                  func(c echo.Context) ([]entity.Transactions, error)
	getTransactionsByMethod          func(c echo.Context, method string) ([]entity.Transactions, error)
	getHistory                       func(c echo.Context, ID string) ([]entity.Transactions, error)
	getHistoryByMethod               func(c echo.Context, ID string, method string) ([]entity.Transactions, error)
	createTransaction                func(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error)
	getTransactionByID               func(c echo.Context, ID string) (*entity.Transactions, error)
	updateTransaction                func(c echo.Context, ID string, transaction *entity.Transactions) (*entity.Transactions, error)
	deleteTransaction                func(c echo.Context, ID string) error
	getCountTransactions             func(c echo.Context) (*entity.GetTransactionsCountView, error)
	createUser                       func(c echo.Context, user entity.Users) (*entity.Users, error)
	createPoints                     func(c echo.Context, userPoints entity.Points) (*entity.Points, error)
	getUserPoints                    func(c echo.Context, ID string) (*entity.Points, error)
	updateUserPoints                 func(c echo.Context, userPoint *entity.Points) error
	getUsersPagination               func(c echo.Context) ([]entity.Users, error)
	getUserByID                      func(c echo.Context, ID string) (*entity.Users, error)
	updateOneByUserId                func(c echo.Context, user *entity.Users) (*entity.Users, error)
	deleteUserById                   func(c echo.Context, ID string) error
	getCountUsers                    func(c echo.Context) (*entity.GetUserCountView, error)
	deleteUserPointsByUserId         func(c echo.Context, ID string) error
)

type allRepositoryMock struct{}

func (r *allRepositoryMock) GetUserLogin(c echo.Context, user entity.LoginBinding) (*entity.Users, error) {
	return getUserLogin(c, user)
}

func (r *allRepositoryMock) GetUserAuth(c echo.Context, user string) (*entity.Users, error) {
	return getUserAuth(c, user)
}

func (r *allRepositoryMock) GetAdminAuth(c echo.Context, user string) (*entity.Users, error) {
	return getAdminAuth(c, user)
}

func (r *allRepositoryMock) GetAuth(c echo.Context, user string) (*entity.Users, error) {
	return getAuth(c, user)
}

func (r *allRepositoryMock) GetProductByID(c echo.Context, productID string) (*entity.Products, error) {
	return getProductByID(c, productID)
}

func (r *allRepositoryMock) GetSerialNumbers(c echo.Context) ([]entity.SerialNumbers, error) {
	return getSerialNumbers(c)
}

func (r *allRepositoryMock) GetSerialNumber(c echo.Context, productID string) (*entity.SerialNumbers, error) {
	return getSerialNumber(c, productID)
}

func (r *allRepositoryMock) UpdateSerialStatus(c echo.Context, serial int64, status string) error {
	return updateSerialStatus(c, serial, status)
}

func (r *allRepositoryMock) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	return createProduct(c, product)
}

func (r *allRepositoryMock) CreateSerialNumber(c echo.Context, serialNumber *entity.SerialNumbers) error {
	return createSerialNumber(c, serialNumber)
}

func (r *allRepositoryMock) GetProducts(c echo.Context) ([]entity.Products, error) {
	return getProducts(c)
}

func (r *allRepositoryMock) GetProductsByMethod(c echo.Context, paymentMethod string) ([]entity.Products, error) {
	return getProductsByMethod(c, paymentMethod)
}

func (r *allRepositoryMock) GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error) {
	return getProductsByCategory(c, category)
}

func (r *allRepositoryMock) UpdateCategoryProduct(c echo.Context, ID string, category string) error {
	return updateCategoryProduct(c, ID, category)
}

func (r *allRepositoryMock) UpdateMethodProduct(c echo.Context, ID string, changeMethodBuy bool) error {
	return updateMethodProduct(c, ID, changeMethodBuy)
}

func (r *allRepositoryMock) UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error) {
	return updateProduct(c, ID, product)
}

func (r *allRepositoryMock) DeleteProduct(c echo.Context, ID string) error {
	return deleteProduct(c, ID)
}

func (r *allRepositoryMock) DeleteAllSerialNumberByProductID(c echo.Context, ID string) error {
	return deleteAllSerialNumberByProductID(c, ID)
}

func (r *allRepositoryMock) DeleteNSerialNumberByProductID(c echo.Context, ID string, N int) error {
	return deleteNSerialNumberByProductID(c, ID, N)
}

func (r *allRepositoryMock) GetUserByIDRaw(c echo.Context, ID string) (entity.Users, error) {
	return getUserByIDRaw(c, ID)
}

func (r *allRepositoryMock) GetProductByIDRaw(c echo.Context, ID string) (entity.Products, error) {
	return getProductByIDRaw(c, ID)
}

func (r *allRepositoryMock) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	return getTransactions(c)
}

func (r *allRepositoryMock) GetTransactionsByMethod(c echo.Context, method string) ([]entity.Transactions, error) {
	return getTransactionsByMethod(c, method)
}

func (r *allRepositoryMock) GetHistory(c echo.Context, ID string) ([]entity.Transactions, error) {
	return getHistory(c, ID)
}

func (r *allRepositoryMock) GetHistoryByMethod(c echo.Context, ID string, method string) ([]entity.Transactions, error) {
	return getHistoryByMethod(c, ID, method)
}

func (r *allRepositoryMock) CreateTransaction(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error) {
	return createTransaction(c, transaction)
}

func (r *allRepositoryMock) GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error) {
	return getTransactionByID(c, ID)
}

func (r *allRepositoryMock) UpdateTransaction(c echo.Context, ID string, transaction *entity.Transactions) (*entity.Transactions, error) {
	return updateTransaction(c, ID, transaction)
}

func (r *allRepositoryMock) DeleteTransaction(c echo.Context, ID string) error {
	return deleteTransaction(c, ID)
}

func (r *allRepositoryMock) GetCountTransactions(c echo.Context) (*entity.GetTransactionsCountView, error) {
	return getCountTransactions(c)
}

func (r *allRepositoryMock) CreateUser(c echo.Context, user entity.Users) (*entity.Users, error) {
	return createUser(c, user)
}

func (r *allRepositoryMock) CreatePoints(c echo.Context, userPoints entity.Points) (*entity.Points, error) {
	return createPoints(c, userPoints)
}

func (r *allRepositoryMock) GetUserPoints(c echo.Context, ID string) (*entity.Points, error) {
	return getUserPoints(c, ID)
}

func (r *allRepositoryMock) UpdateUserPoints(c echo.Context, userPoint *entity.Points) error {
	return updateUserPoints(c, userPoint)
}

func (r *allRepositoryMock) GetUsersPagination(c echo.Context) ([]entity.Users, error) {
	return getUsersPagination(c)
}

func (r *allRepositoryMock) GetUserByID(c echo.Context, ID string) (*entity.Users, error) {
	return getUserByID(c, ID)
}

func (r *allRepositoryMock) UpdateOneByUserId(c echo.Context, user *entity.Users) (*entity.Users, error) {
	return updateOneByUserId(c, user)
}

func (r *allRepositoryMock) DeleteUserById(c echo.Context, ID string) error {
	return deleteUserById(c, ID)
}

func (r *allRepositoryMock) GetCountUsers(c echo.Context) (*entity.GetUserCountView, error) {
	return getCountUsers(c)
}

func (r *allRepositoryMock) DeleteUserPointsByUserId(c echo.Context, ID string) error {
	return deleteUserPointsByUserId(c, ID)
}

func TestGetUserLogin(t *testing.T) {
	// cant work because no compare in mock
	serviceRepo := &allRepositoryMock{}
	service := NewService(serviceRepo)
	ctx := echo.New().NewContext(nil, nil)

	param := entity.LoginBinding{
		Email:    "adminweb@gmail.com",
		Password: "admin567",
	}

	expectedVal := entity.Users{
		Role:     "admin",
		Username: "adminweb",
		Email:    "adminweb@gmail.com",
		Password: "admin567",
	}

	getUserLogin = func(c echo.Context, param entity.LoginBinding) (*entity.Users, error) {
		return &expectedVal, nil
	}

	result, err := service.Login(ctx, param)
	assert.Nil(t, err)

	assert.NotNil(t, result)
	assert.EqualValues(t, expectedVal.Username, result.Username)
	assert.EqualValues(t, expectedVal.Email, result.Email)
	assert.EqualValues(t, expectedVal.Password, result.Password)
}
