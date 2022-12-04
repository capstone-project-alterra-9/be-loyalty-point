package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"

	"errors"
	"github.com/labstack/echo/v4"
)

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