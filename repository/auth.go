package repository

import (
	"capstone-project/entity"
	"errors"

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
	var users []entity.Users
	err := r.connection.Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		if user.Email == u.Email {
			err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
			if err != nil {
				return nil, err
			}
			return &u, nil
		}
	}

	return nil, errors.New("wrong email or password")
}
