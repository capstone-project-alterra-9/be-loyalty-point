package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context, user entity.Users) (*entity.RegisterView, error) {
	user.Points = 20000
	user.ID = (guuid.New()).String();
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

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

	token,_ := jwtAuth.CreateToken(result.Username, result.Email)
	refreshToken,_ := jwtAuth.CreateRefreshToken(result.Username, result.Email)

	return &entity.LoginView{
		Username: 		result.Username,
		Email:    		result.Email,
		Password: 		result.Password,
		Token: 			token,
		RefreshToken: 	refreshToken,
	}, nil
}