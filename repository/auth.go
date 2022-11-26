package repository

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func (r *repository) GetUserLogin(c echo.Context, user entity.LoginBinding) (*entity.Users, error) {
	var userDomain *entity.Users
	err := r.connection.First(&userDomain, "email = ?", user.Email).Error
	if userDomain.ID == "" {
		return nil, err
	}

	return userDomain, nil
}
