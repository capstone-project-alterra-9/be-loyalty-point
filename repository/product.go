package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetProduct(c echo.Context) ([]entity.Products, error) {
	var products []entity.Products
	err := r.connection.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) GetProductByID(c echo.Context, id uint) (entity.Products, error) {
	var product entity.Products
	err := r.connection.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
