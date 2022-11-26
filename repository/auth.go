package repository

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (r *repository) Register(c echo.Context, user entity.Users) (*entity.Users, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Login(c echo.Context, user entity.Users) (*entity.Users, error) {
	var userDomain *entity.Users;
	err := r.connection.First(&userDomain, "email = ?", user.Email).Error

	if userDomain.ID == "" {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(user.Password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}