package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"

	"errors"
	"github.com/labstack/echo/v4"
)

<<<<<<< HEAD
func (s *Service) DeleteOneById(c echo.Context, transactionID string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactionDomain, err := s.repo.GetTransactionByID(c, transactionID)
		if transactionDomain != nil {
			err = s.repo.DeleteOneById(c, transactionID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}
=======
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
>>>>>>> 344fb1ef0bd0d3eccd4bb333f6b6d2d64a329e93
