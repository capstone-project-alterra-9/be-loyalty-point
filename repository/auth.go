package repository

import (
	// "fmt"

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
	err := r.connection.First(&user, "email = ?", user.Email).Error

	// if userDomain.ID == 0 {
	// 	fmt.Println("User not found")
	// 	return nil, err
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}