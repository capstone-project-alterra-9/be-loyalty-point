package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"
	"github.com/labstack/echo/v4"
)

func (s *Service) CreateTransaction(c echo.Context, transaction entity.Transactions) (*entity.Transactions, error) {
	result, err := s.repo.GetUserByUsername(c, middleware.ExtractTokenUsername(c))
	if err != nil {
		return nil, err
	}

	transaction.UserID = result.ID
	transaction.Status = "Pending"
	return s.repo.CreateTransaction(c, transaction)
}
