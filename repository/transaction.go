package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsByCategories(method string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "paymentMethod = ?", method).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsByUser(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
