package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetUserSoftDeleteByID(c echo.Context, ID string) (*entity.Users, error) {
	var userDomain *entity.Users
	err := r.connection.Unscoped().First(&userDomain, "id = ?", ID).Error
	if userDomain.ID == "" {
		return nil, err
	}
	return userDomain, nil
}

func (r *repository) GetProductSoftDeleteByID(c echo.Context, ID string) (*entity.Products, error) {
	var productDomain *entity.Products
	err := r.connection.Unscoped().First(&productDomain, "id = ?", ID).Error
	if productDomain.ID == "" {
		return nil, err
	}
	return productDomain, nil
}

func (r *repository) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsByMethod(c echo.Context, method string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "payment_method = ?", method).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetHistory(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetHistoryByMethod(c echo.Context, ID string, method string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ? AND payment_method = ?", ID, method).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) CreateTransaction(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error) {
	err := r.connection.Create(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *repository) GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error) {
	var transactionDomain *entity.Transactions
	err := r.connection.First(&transactionDomain, "id = ?", ID).Error
	if transactionDomain.ID == "" {
		return nil, err
	}
	return transactionDomain, nil
}

func (r *repository) UpdateTransaction(c echo.Context, ID string, transaction *entity.Transactions) (*entity.Transactions, error) {
	err := r.connection.Where("id = ?", ID).Updates(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *repository) DeleteTransaction(c echo.Context, ID string) error {
	err := r.connection.Where("id = ?", ID).Delete(&entity.Transactions{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCountTransactions(c echo.Context) (*entity.GetTransactionsCountView, error) {
	var transactions []entity.Transactions
	var failedTransactions []entity.Transactions
	var onProgressTransactions []entity.Transactions
	var successTransactions []entity.Transactions

	err := r.connection.Find(&transactions).Error
	err1 := r.connection.Find(&failedTransactions, "status = ?", "failed").Error
	err2 := r.connection.Find(&onProgressTransactions, "status = ?", "pending").Error
	err3 := r.connection.Find(&successTransactions, "status = ?", "success").Error

	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}

	return &entity.GetTransactionsCountView{
		TotalTransactions:           len(transactions),
		TotalFailedTransactions:     len(failedTransactions),
		TotalOnProgressTransactions: len(onProgressTransactions),
		TotalSuccessTransactions:    len(successTransactions),
	}, nil
}
