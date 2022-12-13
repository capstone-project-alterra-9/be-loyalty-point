package service

import (
	"capstone-project/entity"
	"capstone-project/constant"
	"capstone-project/helper"
	jwtAuth "capstone-project/middleware"

	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context, user entity.RegisterBinding) (*entity.RegisterView, error) {
	var userDomain entity.Users
	userDomain.ID = (guuid.New()).String()
	userDomain.Role = constant.USER_TYPE_USER
	var flag bool
	flag = helper.ValidateLength(user.Username, 8, 16)
	if !flag {
		return nil, helper.ErrUsernameLength
	}
	flag = helper.ValidateAlphanumeric(user.Username)
	if !flag {
		return nil, helper.ErrUsernameAlphanumeric
	}

	userDomain.Username = user.Username
	subEmail, err := helper.ValidateEmail(user.Email)
	if err != nil {
		return nil, err
	}
	userDomain.Email = user.Email
	if subEmail == user.Username {
		return nil, helper.ErrEmailUsername
	}

	if len(user.Password) < 8 {
		return nil, helper.ErrPasswordLength
	}
	flag = helper.ValidateAlphanumeric(user.Password)
	if !flag {
		return nil, helper.ErrPasswordAlphanumeric
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	userDomain.Password = string(password)

	result, err := s.repo.CreateUser(c, userDomain)
	if err != nil {
		return nil, err
	}

	userPoints := entity.Points{
		ID:         (guuid.New()).String(),
		UserID:     result.ID,
		Points:     20000,
		CostPoints: 0,
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
		ID:			  result.ID,
		Username:     result.Username,
		Email:        result.Email,
		Password:     result.Password,
		Token:        token,
		RefreshToken: refreshToken,
		Account:      result.Role,
	}, nil
}

func (s *Service) ReGenerateToken(c echo.Context, refreshToken entity.TokenBinding) (*entity.RefreshTokenView, error) {
	
	user, err := jwtAuth.ValidateToken(refreshToken)
	
	if err != nil {
		return nil, err
	}

	token, _ := jwtAuth.CreateToken(user.Username, user.Email)

	return &entity.RefreshTokenView{
		Token:        token,
		Duration: 	"ini duration",
	}, nil
}
