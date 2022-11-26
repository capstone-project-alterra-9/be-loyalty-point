package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	var (
		user   = middleware.ExtractTokenUsername(c)
		admins *entity.Admins
		users  *entity.Users
		err    error
	)
	// if user an admin
	admins, err = s.repo.GetAdminAuth(c, user)
	if admins != nil {
		transactions, err := s.repo.GetTransactions(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}

	// if user a customer
	users, err = s.repo.GetUserAuth(c, user)
	if users != nil {
		transactions, err := s.repo.GetTransactionsByUser(c, users.ID)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}

	return nil, err
}
