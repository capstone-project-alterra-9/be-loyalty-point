package repository

import (
	"capstone-project/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions).Error
	if err != nil {
		return nil, errors.New("no transactions found")
	}
	return transactions, nil
}

func (r *repository) GetTransactionsByUser(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Where("id = ?", ID).Find(&transactions).Error
	if err != nil {
		return nil, errors.New("no transactions found")
	}

	return transactions, nil
}
