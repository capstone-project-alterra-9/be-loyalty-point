package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetUserAuth(c echo.Context) ([]entity.Users, error) {
	var users []entity.Users
	err := r.connection.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
