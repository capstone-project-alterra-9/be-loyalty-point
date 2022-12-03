package service

import (
	"capstone-project/entity"
	// jwtAuth "capstone-project/middleware"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetUsersPagination(c echo.Context) ([]entity.Users, error) {
	// user := jwtAuth.ExtractTokenUsername(c)
	// adminAuth, err := s.repo.GetAdminAuth(c, user)
	// if adminAuth != nil {
		users, err := s.repo.GetUsersPagination(c)
		if err != nil {
			return nil, err
		}
		return users, nil
	// }
	// return nil, err
}
