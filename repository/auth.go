package repository

import (
	"capstone-project/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetAdminAuth(c echo.Context) ([]entity.Admins, error) {
	var admins []entity.Admins
	err := r.connection.Find(&admins).Error
	if err != nil {
		return nil, errors.New("you are not an admin")
	}
	return admins, nil
}
