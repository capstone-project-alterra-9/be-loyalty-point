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

func (r *repository) GetUserAuth(c echo.Context, user string) (*entity.Users, error) {
	var userDomain *entity.Users
	err := r.connection.First(&userDomain, "username = ?", user).Error
	if userDomain.ID == "" {
		return nil, err
	}

	return userDomain, nil
}

func (r *repository) GetAdminAuth(c echo.Context, user string) (*entity.Admins, error) {
	var adminDomain *entity.Admins
	err := r.connection.First(&adminDomain, "username = ?", user).Error
	if adminDomain.ID == "" {
		return nil, err
	}

	return adminDomain, nil
}