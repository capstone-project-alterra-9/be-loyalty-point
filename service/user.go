package service

import (
	"capstone-project/entity"
	"capstone-project/helper"
	jwtAuth "capstone-project/middleware"

	"errors"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// buat fitur creat by admin, kurleb sama kaya register -- uda
// update user pake data yang sama --
// ada bug di update point

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
				Points:     userPoint.Points,
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
		tempCompare := userData
		userPoint, err := s.repo.GetUserPoints(c, userData.ID)
		if err != nil {
			return nil, err
		}

		if user.Role != "" {
			if user.Role == "admin" || user.Role == "user" {
				userData.Role = user.Role
			} else {
				return nil, errors.New("role must be admin or user")
			}
		}

		var flag bool
		if user.Username != "" {
			flag = helper.ValidateLength(user.Username, 8, 16)
			if !flag {
				return nil, helper.ErrUsernameLength
			}
			flag = helper.ValidateAlphanumeric(user.Username)
			if !flag {
				return nil, helper.ErrUsernameAlphanumeric
			}
			userData.Username = user.Username
		}

		if user.Email != "" {
			subEmail, err := helper.ValidateEmail(user.Email)
			if err != nil {
				return nil, err
			}
			userData.Email = user.Email
			if subEmail == user.Username {
				return nil, helper.ErrEmailUsername
			}
		}

		if user.Password != "" {
			if len(user.Password) < 8 {
				return nil, helper.ErrPasswordLength
			}
			flag = helper.ValidateAlphanumeric(user.Password)
			if !flag {
				return nil, helper.ErrPasswordAlphanumeric
			}
			password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if string(password) != userData.Password {
				userData.Password = string(password)
			}
		}

		if user.Points != userPoint.Points || user.CostPoints != userPoint.CostPoints {
			userPoint.Points = user.Points
			userPoint.CostPoints = user.CostPoints
			_ = s.repo.UpdateUserPoints(c, userPoint)
			if err != nil {
				return nil, err
			}
		}

		if userData == tempCompare && user.Points == userPoint.Points && user.CostPoints == userPoint.CostPoints {
			return nil, helper.ErrSameDataRequest
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
			Points:     userPoint.Points,
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

func (s *Service) GetForgotPassword(c echo.Context) error {
	// user := jwtAuth.ExtractTokenUsername(c)
	
	m := gomail.NewMessage()
	m.SetHeader("From", "aji.zapar00@gmail.com")
	m.SetHeader("To", "ajizapar080500@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "aji.zapar00@gmail.com", "brrjidfupjwvyvta")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}