package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"
	"errors"

	guuid "github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactions, err := s.repo.GetTransactions(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsRedeem(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactions, err := s.repo.GetTransactionsRedeem(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsBuy(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactions, err := s.repo.GetTransactionsBuy(c)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		transactions, err := s.repo.GetTransactionByID(c, ID)
		if err != nil {
			return nil, err
		}
		if (auth.Role == "user" && auth.ID == transactions.UserID) || auth.Role == "admin" {
			return transactions, nil
		} else {
			return nil, errors.New("unauthorized")
		}
	}
	return nil, err
}

func (s *Service) GetHistory(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		transactions, err := s.repo.GetHistory(c, userDomain.ID)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetHistoryBuy(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		transactions, err := s.repo.GetHistoryBuy(c, userDomain.ID)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) GetHistoryRedeem(c echo.Context) ([]entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		transactions, err := s.repo.GetHistoryRedeem(c, userDomain.ID)
		if err != nil {
			return nil, err
		}
		return transactions, nil
	}
	return nil, err
}

func (s *Service) CreateTransactionByUser(c echo.Context, transaction entity.TransactionsBinding) (*entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		product, err := s.repo.GetProduct(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		userPoints, err := s.repo.GetUserPoints(c, userDomain.ID)
		if err != nil {
			return nil, err
		}
		serial, err := s.repo.GetSerialNumber(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		transactionDomain := &entity.Transactions{
			ID:            (guuid.New()).String(),
			PaymentMethod: transaction.PaymentMethod,
			UserID:        userDomain.ID,
			ProductID:     transaction.ProductID,
			SerialNumber:  serial.Serial,
			IdentifierNum: transaction.IdentifierNum,
			Price:         product.Price,
			Status:        "pending",
		}

		if transaction.PaymentMethod == "redeem" && userPoints.Points < product.Price {
			transactionDomain.Status = "failed"
		} else if transaction.PaymentMethod == "buy" {
			// midtrans payment gateway logic
			// if midtrans payment gateway succes then update user points ( increase cost points )
			return nil, errors.New("midtrans payment gateway not implemented yet")
		}

		result, err := s.repo.CreateTransaction(c, transactionDomain)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, err
}

func (s *Service) CreateTransactionByAdmin(c echo.Context, transaction entity.Transactions) (*entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transaction.ID = (guuid.New()).String() + "-dummy"
		product, err := s.repo.GetProduct(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		transaction.Price = product.Price
		serial, err := s.repo.GetSerialNumber(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		transaction.SerialNumber = serial.Serial
		transaction.Status = "pending"
		result, err := s.repo.CreateTransaction(c, &transaction)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, err
}

func (s *Service) UpdateTransactionByAdmin(c echo.Context, ID string, transaction entity.UpdateTransactionBinding) (*entity.Transactions, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactionDomain, err := s.repo.GetTransactionByID(c, ID)
		if err != nil {
			return nil, err
		}
		if transaction.Status == "succes" || transaction.Status == "Succes" {
			if transactionDomain.PaymentMethod == "redeem" {
				err = s.repo.UpdateSerialStatus(c, transactionDomain.SerialNumber, "unavailable")
				if err != nil {
					return nil, err
				}
				userPoint, err := s.repo.GetUserPoints(c, transactionDomain.UserID)
				if err != nil {
					return nil, err
				}
				userPoint.Points = userPoint.Points - transactionDomain.Price
				userPoint.CostPoints = userPoint.CostPoints + transactionDomain.Price
				err = s.repo.UpdateUserPoints(c, userPoint)
				if err != nil {
					return nil, err
				}
				transactionDomain.Status = transaction.Status
				result, err := s.repo.UpdateTransaction(c, transactionDomain)
				if err != nil {
					return nil, err
				}
				return result, nil
			} else if transactionDomain.PaymentMethod == "buy" {
				// midtrans payment gateway logic
				return nil, errors.New("midtrans payment gateway not implemented yet")
			}
		} else if transaction.Status == "failed" || transaction.Status == "Failed" {
			transactionDomain.Status = transaction.Status
			result, err := s.repo.UpdateTransaction(c, transactionDomain)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, errors.New("invalid status")
		}
	}
	return nil, err
}

func (s *Service) DeleteTransactionByAdmin(c echo.Context, transactionID string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		err := s.repo.DeleteTransaction(c, transactionID)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
