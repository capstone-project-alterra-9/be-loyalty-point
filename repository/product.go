package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetProductByID(c echo.Context, productID string) (*entity.Products, error) {
	var productsDomain *entity.Products
	err := r.connection.First(&productsDomain, "id = ?", productID).Error
	if productsDomain.ID == "" {
		return nil, err
	}
	return productsDomain, nil
}

func (r *repository) GetSerialNumbers(c echo.Context) ([]entity.SerialNumbers, error) {
	var serialNumbers []entity.SerialNumbers
	err := r.connection.Find(&serialNumbers).Error
	return serialNumbers, err
}

func (r *repository) GetSerialNumber(c echo.Context, productID string) (*entity.SerialNumbers, error) {
	var serialNumber *entity.SerialNumbers
	err := r.connection.First(&serialNumber, "product_id = ? AND status = ?", productID, "available").Error
	if serialNumber.ID == "" {
		return nil, err
	}
	return serialNumber, nil
}

func (r *repository) UpdateSerialStatus(c echo.Context, serial int64, status string) error {
	err := r.connection.Model(&entity.SerialNumbers{}).Where("serial = ?", serial).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	err := r.connection.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) CreateSerialNumber(c echo.Context, serialNumber *entity.SerialNumbers) error {
	err := r.connection.Create(&serialNumber).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetProducts(c echo.Context) ([]entity.Products, error) {
	var products []entity.Products
	err := r.connection.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error) {
	var products []entity.Products
	err := r.connection.Find(&products, "category = ?", category).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error) {
	err := r.connection.Model(&entity.Products{}).Where("id = ?", ID).Updates(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repository) DeleteProduct(c echo.Context, ID string) error {
	err := r.connection.Where("id = ?", ID).Delete(&entity.Products{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteAllSerialNumberByProductID(c echo.Context, ID string) error {
	err := r.connection.Where("product_id = ?", ID).Delete(&entity.SerialNumbers{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteNSerialNumberByProductID(c echo.Context, ID string, N int) error {
	err := r.connection.Where("product_id = ?", ID).Limit(N).Delete(&entity.SerialNumbers{}).Error
	if err != nil {
		return err
	}
	return nil
}
