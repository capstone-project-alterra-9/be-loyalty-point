package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) CreateUser(c echo.Context, user entity.Users) (*entity.Users, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreatePoints(c echo.Context, userPoints entity.Points) (*entity.Points, error) {
	err := r.connection.Create(&userPoints).Error
	if err != nil {
		return nil, err
	}
	return &userPoints, nil
}

func (r *repository) GetUserPoints(c echo.Context, ID string) (*entity.Points, error) {
	var userPoints *entity.Points
	err := r.connection.First(&userPoints, "user_id = ?", ID).Error
	if userPoints.ID == "" {
		return nil, err
	}
	return userPoints, nil
}

func (r *repository) UpdateUserPoints(c echo.Context, userPoint *entity.Points) error {
	err := r.connection.Where("user_id = ?", userPoint.UserID).Updates(userPoint).Error
	if err != nil {
		return err
	}
	return nil
}
