package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"

	"errors"
	"github.com/labstack/echo/v4"
)

func (s *Service) DeleteOneById(c echo.Context, userId string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)

	if adminAuth != nil {
		userData, err := s.repo.GetUserByID(c, userId)
		if userData != nil {
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

func (s *Service) GetUserById(c echo.Context, ID string) (*entity.Users, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		userData, err := s.repo.GetUserByID(c, ID)
		if err != nil {
			return nil, err
		}
		if (auth.Role == "user" && auth.ID == userData.ID) || auth.Role == "admin" {
			return userData, nil
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


func (s *Service) UpdateOneById(c echo.Context, ID string, user entity.Users) (*entity.Users, error) {
	userAuth := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, userAuth)
	if adminAuth != nil {
		userData, err := s.repo.GetUserByID(c, ID)
		if err != nil {
			return nil, err
		}

		userData.Username = user.Username;
		userData.Email = user.Email; 
		userData.Password = user.Password;

		result, err := s.repo.UpdateOneByUserId(c, userData)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, err
}