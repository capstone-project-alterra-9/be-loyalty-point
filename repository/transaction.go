package repository

import (
	"capstone-project/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func (r *repository) CreateTransaction(c echo.Context, transaction entity.Transactions) (entity.Transactions, error) {
	err := r.connection.Create(&transaction).Error
	if err != nil {
		return transaction, errors.New("failed to create transaction")
	}
	return transaction, nil
}

func (r *repository) GetSerialByStatus(c echo.Context, status string) (entity.SerialNumbers, error) {
	var serial entity.SerialNumbers
	err := r.connection.Where("status = ?", status).First(&serial).Error
	if err != nil {
		return serial, errors.New("serial with status " + status + " not found")
	}
	return serial, nil
}

func (r *repository) UpdateSerialStatus(c echo.Context, updateResult entity.SerialNumbers) error {
	err := r.connection.Save(&updateResult).Error
	if err != nil {
		return errors.New("failed to update serial status")
	}
	return nil
}
