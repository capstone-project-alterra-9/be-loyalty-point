package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"

	"errors"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) DeleteOneById(c echo.Context, userId string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)

	if adminAuth != nil {
		userData, err := s.repo.GetUserByID(c, userId)
		if userData != nil {
			err = s.repo.DeleteUserPointsByUserId(c, userId)
			if err != nil {
				return err
			}
			err = s.repo.DeleteUserById(c, userId)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}

func (s *Service) GetUserById(c echo.Context, ID string) (*entity.UsersView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		userData, err := s.repo.GetUserByID(c, ID)
		if err != nil {
			return nil, err
		}
		if (auth.Role == "user" && auth.ID == userData.ID) || auth.Role == "admin" {
			userPoint, err := s.repo.GetUserPoints(c, userData.ID)
			if err != nil {
				return nil, err
			}
			return &entity.UsersView{
				ID:         userData.ID,
				CreatedAt:  userData.CreatedAt,
				UpdatedAt:  userData.UpdatedAt,
				Role:       userData.Role,
				Username:   userData.Username,
				Email:      userData.Email,
				Password:   userData.Password,
				Point:      userPoint.Points,
				CostPoints: userPoint.CostPoints,
			}, nil
		} else {
			return nil, errors.New("unauthorized")
		}
	}
	return nil, err
}

func (s *Service) GetUsersPagination(c echo.Context) ([]entity.Users, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		users, err := s.repo.GetUsersPagination(c)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	return nil, err
}

func (s *Service) UpdateOneById(c echo.Context, ID string, user entity.UpdateUserBinding) (*entity.UsersView, error) {
	userAuth := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, userAuth)
	if adminAuth != nil {
		userData, err := s.repo.GetUserByID(c, ID)
		if err != nil {
			return nil, err
		}

		if user.Role != "" {
			if user.Role == "admin" || user.Role == "user" {
				userData.Role = user.Role
			} else {
				return nil, errors.New("Role must be admin or user")
			}
		}

		userData.Username = user.Username
		userData.Email = user.Email
		userData.Password = user.Password

		userPoint, err := s.repo.GetUserPoints(c, userData.ID)
		if err != nil {
			return nil, err
		}

		if user.Point != userPoint.Points || user.CostPoints != userPoint.CostPoints {
			userPoint.Points = user.Point
			userPoint.CostPoints = user.CostPoints
			_ = s.repo.UpdateUserPoints(c, userPoint)
			if err != nil {
				return nil, err
			}
		}

		updatedUser, err := s.repo.UpdateOneByUserId(c, userData)
		if err != nil {
			return nil, err
		}
		return &entity.UsersView{
			ID:         updatedUser.ID,
			CreatedAt:  updatedUser.CreatedAt,
			UpdatedAt:  updatedUser.UpdatedAt,
			Role:       updatedUser.Role,
			Username:   updatedUser.Username,
			Email:      updatedUser.Email,
			Password:   updatedUser.Password,
			Point:      userPoint.Points,
			CostPoints: userPoint.CostPoints,
		}, nil
	}
	return nil, err
}

func (s *Service) CreateUserByAdmin(c echo.Context, user entity.CreateUserBinding) (*entity.CreateUserView, error) {
	userAuth := jwtAuth.ExtractTokenUsername(c)
	adminAuth, _ := s.repo.GetAdminAuth(c, userAuth)

	if adminAuth == nil {
		return nil, errors.New("Unauthorized")
	}

	var userDomain entity.Users
	userDomain.ID = (guuid.New()).String()
	userDomain.Role = "user"
	userDomain.Username = user.Username
	userDomain.Email = user.Email
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	userDomain.Password = string(password)

	result, err := s.repo.CreateUser(c, userDomain)
	if err != nil {
		return nil, err
	}

	userPoint := entity.Points{
		ID:         (guuid.New()).String(),
		UserID:     result.ID,
		Points:     user.Point,
		CostPoints: 0,
	}

	_, err = s.repo.CreatePoints(c, userPoint)
	if err != nil {
		return nil, err
	}

	return &entity.CreateUserView{
		ID:         result.ID,
		Username:   result.Username,
		Email:      result.Email,
		Password:   result.Password,
		Point:      userPoint.Points,
		CostPoints: userPoint.CostPoints,
	}, nil
}

func (s *Service) GetCountUsers(c echo.Context) (*entity.GetUserCountView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, _ := s.repo.GetAdminAuth(c, user)
	if adminAuth == nil {
		return nil, errors.New("Unauthorized")
	}

	userCount, err := s.repo.GetCountUsers(c)

	if err != nil {
		return nil, err
	}
	return userCount, nil
}
