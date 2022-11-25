package service

import (
	"capstone-project/entity"
	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context, user entity.Users) (*entity.Users, error) {
	user.Points = 20000
	return s.repo.Register(c, user)
}
