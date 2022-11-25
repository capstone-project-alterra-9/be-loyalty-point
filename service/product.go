package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"
	"errors"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	result, err := s.repo.GetAdmins(c)
	if err != nil {
		return nil, err
	}
	for _, admin := range result {
		if admin.Username == middleware.ExtractTokenUsername(c) {
			product.CreatedBy = middleware.ExtractTokenUsername(c)
			product.UpdatedBy = product.CreatedBy
			return s.repo.CreateProduct(c, product)
		}
	}
	return nil, errors.New("You are not an admin")
}
