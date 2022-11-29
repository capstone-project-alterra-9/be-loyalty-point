package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetProduct(c echo.Context, productID string) (*entity.Products, error) {
	var productsDomain *entity.Products
	err := r.connection.First(&productsDomain, "id = ?", productID).Error
	if productsDomain.ID == "" {
		return nil, err
	}
	return productsDomain, nil
}

func (r *repository) GetSerialNumber(c echo.Context, productID string) (*entity.SerialNumbers, error) {
	var serialNumber *entity.SerialNumbers
	err := r.connection.First(&serialNumber, "product_id = ? AND status = ?", productID, "available").Error
	if serialNumber.ID == "" {
		return nil, err
	}
	return serialNumber, nil
}

func (r *repository) UpdateSerialStatus(c echo.Context, serial uint64, status string) error {
	err := r.connection.Model(&entity.SerialNumbers{}).Where("serial = ?", serial).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
