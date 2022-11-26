package service

import (
	"capstone-project/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetTransactions(c echo.Context, user string) ([]entity.Transactions, error) {
	var (
		admins []entity.Admins
		users  []entity.Users
		err    error
	)
	// if user an admin
	admins, err = s.repo.GetAdminAuth(c)
	if err != nil {
		return nil, err
	}
	for _, admin := range admins {
		if user == admin.Username {
			transactions, err := s.repo.GetTransactions(c)
			if err != nil {
				return nil, err
			}
			return transactions, nil
		}
	}

	// if user a customer
	users, err = s.repo.GetUserAuth(c)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		if user == u.Username {
			transactions, err := s.repo.GetTransactionsByUser(c, u.ID)
			if err != nil {
				return nil, err
			}
			return transactions, nil
		}
	}

	return nil, errors.New("you are not authorized")
}
