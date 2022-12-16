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

func (r *repository) GetUserPointIgnoreEmpty(c echo.Context, ID string) (*entity.Points, error) {
	var userPoints *entity.Points
	err := r.connection.Find(&userPoints, "user_id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return userPoints, nil
}

func (r *repository) UpdateUserPoints(c echo.Context, userPoint *entity.Points) error {
	err := r.connection.Model(&entity.Points{}).Where("user_id = ?", userPoint.UserID).Updates(map[string]interface{}{
		"user_id":     userPoint.UserID,
		"points":      userPoint.Points,
		"cost_points": userPoint.CostPoints,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUsersPagination(c echo.Context) ([]entity.Users, error) {
	var users []entity.Users

	err := r.connection.Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetUserByID(c echo.Context, ID string) (*entity.Users, error) {
	var userData *entity.Users
	err := r.connection.First(&userData, "id = ?", ID).Error
	if userData.ID == "" {
		return nil, err
	}
	return userData, nil
}

func (r *repository) DeleteUserById(c echo.Context, ID string) error {
	err := r.connection.Delete(&entity.Users{}, "id = ?", ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateOneByUserId(c echo.Context, user *entity.Users) (*entity.Users, error) {
	err := r.connection.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetCountUsers(c echo.Context) (*entity.GetUserCountView, error) {
	var users []entity.Users

	err := r.connection.Find(&users, "role = ?", "user").Error

	if err != nil {
		return nil, err
	}
	return &entity.GetUserCountView{
		TotalCount: len(users),
	}, nil
}

func (r *repository) DeleteUserPointsByUserId(c echo.Context, ID string) error {
	err := r.connection.Where("user_id = ?", ID).Delete(&entity.Points{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserByEmail(c echo.Context, email string) (*entity.Users, error) {
	var userDomain *entity.Users
	err := r.connection.First(&userDomain, "email = ?", email).Error
	if userDomain.ID == "" {
		return nil, err
	}
	return userDomain, nil
}