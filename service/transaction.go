package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"
	"errors"
	"github.com/labstack/echo/v4"
)

func (s *Service) CreateTransaction(c echo.Context, param entity.TransactionBinding) (entity.Transactions, error) {
	var transaction entity.Transactions
	users, err := s.repo.GetUserAuth(c)
	if err != nil {
		return transaction, err
	}

	for _, user := range users {
		if user.Username == middleware.ExtractTokenUsername(c) {
			transaction.UserID = user.ID
			transaction.ProductID = param.ProductID
			var serial entity.SerialNumbers
			serial, err = s.repo.GetSerialByStatus(c, "Available")
			if err != nil {
				return transaction, err
			}
			transaction.SerialNumber = serial.Serial
			var product entity.Products
			product, err = s.repo.GetProductByID(c, transaction.ProductID)
			if err != nil {
				return transaction, err
			}
			transaction.IdentifierNum = param.IdentifierNum
			transaction.TotalPrice = product.Price
			transaction.Status = "Pending"

			var result entity.Transactions
			result, err = s.repo.CreateTransaction(c, transaction)
			if err != nil {
				return transaction, err
			}

			serial.Status = "Unavailable"
			err = s.repo.UpdateSerialStatus(c, serial)
			if err != nil {
				return transaction, err
			}

			return result, nil
		}
	}

	return transaction, errors.New("you are not authorized")
}
