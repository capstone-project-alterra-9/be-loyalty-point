package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context, user entity.RegisterBinding) (*entity.RegisterView, error) {
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

	userPoints := entity.Points{
		ID:     (guuid.New()).String(),
		UserID: result.ID,
		Points: 20000,
	}
	_, err = s.repo.CreatePoints(c, userPoints)
	if err != nil {
		return nil, err
	}

	return &entity.RegisterView{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}

func (s *Service) Login(c echo.Context, user entity.LoginBinding) (*entity.LoginView, error) {
	result, err := s.repo.GetUserLogin(c, user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	token, _ := jwtAuth.CreateToken(result.Username, result.Email)
	refreshToken, _ := jwtAuth.CreateRefreshToken(result.Username, result.Email)

	return &entity.LoginView{
		Username:     result.Username,
		Email:        result.Email,
		Password:     result.Password,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
