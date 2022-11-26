package service

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context, user entity.Users) (*entity.RegisterView, error) {
	user.Points = 20000
	result, err := s.repo.Register(c, user)
	if err != nil {
		return nil, err
	}
	return &entity.RegisterView{
		ID:			result.ID,
		Username: 	result.Username,
		Email:    	result.Email,
		Password: 	result.Password,
	}, nil
}

func (s *Service) Login(c echo.Context, user entity.Users) (*entity.LoginView, error) {

	result, err := s.repo.Login(c, user)
	if err != nil {
		return nil, err
	}

	// token := uu.jwtAuth.GenerateToken(int(user.ID))
	token := "test"

	return &entity.LoginView{
		Username: 	result.Username,
		Email:    	result.Email,
		Password: 	result.Password,
		Token: 		token,
	}, nil
}