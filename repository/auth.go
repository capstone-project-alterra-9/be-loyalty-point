package repository

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func (r *repository) Register(c echo.Context, user entity.Users) (*entity.Users, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
