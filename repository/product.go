package repository

import (
	"capstone-project/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func (r *repository) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	err := r.connection.Create(&product).Error
	if err != nil {
		return nil, errors.New("failed to create product")
	}
	return product, nil
}

func (r *repository) CreateSerialNumber(c echo.Context, serialNumber *entity.SerialNumbers) error {
	err := r.connection.Create(&serialNumber).Error
	if err != nil {
		return errors.New("failed to create serial number")
	}
	return nil
}

func (r *repository) GetSerialNumbers(c echo.Context) ([]entity.SerialNumbers, error) {
	var serialNumbers []entity.SerialNumbers
	err := r.connection.Find(&serialNumbers).Error
	if err != nil {
		return nil, errors.New("no serial number found")
	}
	return serialNumbers, nil
}
