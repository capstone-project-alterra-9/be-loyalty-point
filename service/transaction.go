package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	user := middleware.ExtractTokenUsername(c)
	admins, err := s.repo.GetAdminAuth(c, user)
	if admins != nil {
		transactions, err := s.repo.GetTransactions(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsRedeem(c echo.Context) ([]entity.Transactions, error) {
	user := middleware.ExtractTokenUsername(c)
	admins, err := s.repo.GetAdminAuth(c, user)
	if admins != nil {
		transactions, err := s.repo.GetTransactionsRedeem(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsBuy(c echo.Context) ([]entity.Transactions, error) {
	user := middleware.ExtractTokenUsername(c)
	admins, err := s.repo.GetAdminAuth(c, user)
	if admins != nil {
		transactions, err := s.repo.GetTransactionsBuy(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsByUser(c echo.Context) ([]entity.Transactions, error) {
	user := middleware.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		transactions, err := s.repo.GetTransactionsByUser(c, userDomain.ID)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}
