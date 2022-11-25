package repository

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func (r *repository) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	err := r.connection.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) GetAdmins(c echo.Context) ([]entity.Admins, error) {
	var admins []entity.Admins
	err := r.connection.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}
