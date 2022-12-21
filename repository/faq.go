package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) CreateFaq(c echo.Context, faq entity.FAQ) (*entity.FAQ, error) {
	err := r.connection.Create(&faq).Error
	if err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *repository) GetFaqs(c echo.Context) ([]entity.FAQ, error) {
	var faqs []entity.FAQ

	err := r.connection.Find(&faqs).Error

	if err != nil {
		return nil, err
	}

	return faqs, nil
}