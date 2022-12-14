package service

import (
	"capstone-project/entity"
	"capstone-project/helper"
	jwtAuth "capstone-project/middleware"
	"errors"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (s *Service) GetTransactions(c echo.Context) ([]entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactions, err := s.repo.GetTransactions(c)
		if err != nil {
			return nil, err
		}

		var transactionsView []entity.TransactionsView
		for _, transaction := range transactions {
			product, err := s.repo.GetProductByIDRaw(c, transaction.ProductID)
			if err != nil {
				return nil, err
			}
			user, err := s.repo.GetUserByIDRaw(c, transaction.UserID)
			if err != nil {
				return nil, err
			}
			transactionsView = append(transactionsView, entity.TransactionsView{
				ID:            transaction.ID,
				CreatedAt:     transaction.CreatedAt,
				UpdatedAt:     transaction.UpdatedAt,
				PaymentMethod: transaction.PaymentMethod,
				UserID:        transaction.UserID,
				Username:      user.Username,
				ProductID:     transaction.ProductID,
				ProductName:   product.Name,
				Category:      product.Category,
				SerialNumber:  transaction.SerialNumber,
				IdentifierNum: transaction.IdentifierNum,
				Price:         transaction.Price,
				Status:        transaction.Status,
			})
		}
		return transactionsView, nil
	}
	return nil, err
}

func (s *Service) GetTransactionsByMethod(c echo.Context, method string) ([]entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		if method == "buy" || method == "redeem" {
			transactions, err := s.repo.GetTransactionsByMethod(c, method)
			if err != nil {
				return nil, err
			}

			var transactionsView []entity.TransactionsView
			for _, transaction := range transactions {
				product, err := s.repo.GetProductByIDRaw(c, transaction.ProductID)
				if err != nil {
					return nil, err
				}
				user, err := s.repo.GetUserByIDRaw(c, transaction.UserID)
				if err != nil {
					return nil, err
				}
				transactionsView = append(transactionsView, entity.TransactionsView{
					ID:            transaction.ID,
					CreatedAt:     transaction.CreatedAt,
					UpdatedAt:     transaction.UpdatedAt,
					PaymentMethod: transaction.PaymentMethod,
					UserID:        transaction.UserID,
					Username:      user.Username,
					ProductID:     transaction.ProductID,
					ProductName:   product.Name,
					Category:      product.Category,
					SerialNumber:  transaction.SerialNumber,
					IdentifierNum: transaction.IdentifierNum,
					Price:         transaction.Price,
					Status:        transaction.Status,
				})
			}
			return transactionsView, nil
		} else {
			return nil, errors.New("method not found")
		}
	}
	return nil, err
}

func (s *Service) GetTransactionByID(c echo.Context, ID string) (*entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		transaction, err := s.repo.GetTransactionByID(c, ID)
		if err != nil {
			return nil, err
		}
		if (auth.Role == "user" && auth.ID == transaction.UserID) || auth.Role == "admin" {
			product, err := s.repo.GetProductByIDRaw(c, transaction.ProductID)
			if err != nil {
				return nil, err
			}
			user, err := s.repo.GetUserByIDRaw(c, transaction.UserID)
			if err != nil {
				return nil, err
			}
			return &entity.TransactionsView{
				ID:            transaction.ID,
				CreatedAt:     transaction.CreatedAt,
				UpdatedAt:     transaction.UpdatedAt,
				PaymentMethod: transaction.PaymentMethod,
				UserID:        transaction.UserID,
				Username:      user.Username,
				ProductID:     transaction.ProductID,
				ProductName:   product.Name,
				Category:      product.Category,
				SerialNumber:  transaction.SerialNumber,
				IdentifierNum: transaction.IdentifierNum,
				Price:         transaction.Price,
				Status:        transaction.Status,
			}, nil
		} else {
			return nil, errors.New("unauthorized")
		}
	}
	return nil, err
}

func (s *Service) GetHistory(c echo.Context) ([]entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		transactions, err := s.repo.GetHistory(c, userDomain.ID)
		if err != nil {
			return nil, err
		}

		var transactionsView []entity.TransactionsView
		for _, transaction := range transactions {
			product, err := s.repo.GetProductByIDRaw(c, transaction.ProductID)
			if err != nil {
				return nil, err
			}
			user, err := s.repo.GetUserByIDRaw(c, transaction.UserID)
			if err != nil {
				return nil, err
			}
			transactionsView = append(transactionsView, entity.TransactionsView{
				ID:            transaction.ID,
				CreatedAt:     transaction.CreatedAt,
				UpdatedAt:     transaction.UpdatedAt,
				PaymentMethod: transaction.PaymentMethod,
				UserID:        transaction.UserID,
				Username:      user.Username,
				ProductID:     transaction.ProductID,
				ProductName:   product.Name,
				Category:      product.Category,
				SerialNumber:  transaction.SerialNumber,
				IdentifierNum: transaction.IdentifierNum,
				Price:         transaction.Price,
				Status:        transaction.Status,
			})
		}
		return transactionsView, nil
	}
	return nil, err
}

func (s *Service) GetHistoryByMethod(c echo.Context, method string) ([]entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		if method == "buy" || method == "redeem" {
			transactions, err := s.repo.GetHistoryByMethod(c, userDomain.ID, method)
			if err != nil {
				return nil, err
			}

			var transactionsView []entity.TransactionsView
			for _, transaction := range transactions {
				product, err := s.repo.GetProductByIDRaw(c, transaction.ProductID)
				if err != nil {
					return nil, err
				}
				user, err := s.repo.GetUserByIDRaw(c, transaction.UserID)
				if err != nil {
					return nil, err
				}
				transactionsView = append(transactionsView, entity.TransactionsView{
					ID:            transaction.ID,
					CreatedAt:     transaction.CreatedAt,
					UpdatedAt:     transaction.UpdatedAt,
					PaymentMethod: transaction.PaymentMethod,
					UserID:        transaction.UserID,
					Username:      user.Username,
					ProductID:     transaction.ProductID,
					ProductName:   product.Name,
					Category:      product.Category,
					SerialNumber:  transaction.SerialNumber,
					IdentifierNum: transaction.IdentifierNum,
					Price:         transaction.Price,
					Status:        transaction.Status,
				})
			}
			return transactionsView, nil
		} else {
			return nil, errors.New("method not found")
		}
	}
	return nil, err
}

func (s *Service) GetHistoryByMethodCategory(c echo.Context, method string, category string) ([]entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		if (method == "buy" || method == "redeem") && (category == "credits" || category == "data-quota" || category == "e-money" || category == "cashout") {
			transactions, err := s.GetHistoryByMethod(c, method)
			if err != nil {
				return nil, err
			}

			var result []entity.TransactionsView
			for _, transaction := range transactions {
				if transaction.Category == category {
					result = append(result, transaction)
				}
			}
			return result, nil
		} else {
			return nil, errors.New("method or category not found")
		}
	}
	return nil, err
}

func (s *Service) CreateTransactionByUser(c echo.Context, transaction entity.TransactionsBinding) (*entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	userDomain, err := s.repo.GetUserAuth(c, user)
	if userDomain != nil {
		product, err := s.repo.GetProductByID(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		if transaction.IdentifierNum == "" {
			return nil, errors.New("identifier number is required")
		}
		err = helper.ValidateIdentifierNumber(transaction.IdentifierNum)
		if err != nil {
			return nil, err
		}

		if transaction.PaymentMethod == "buy" && product.Buy || transaction.PaymentMethod == "redeem" && product.Redeem {
			if product.Stock == 0 {
				return nil, errors.New("product out of stock")
			}
			userPoint, err := s.repo.GetUserPoints(c, userDomain.ID)
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

			if transaction.PaymentMethod == "redeem" && userPoint.Points < product.Price {
				transactionDomain.Status = "failed"
				return nil, errors.New("not enough points")
			} else if transaction.PaymentMethod == "redeem" && userPoint.Points >= product.Price {
				transactionDomain.Status = "success"
				userPoint.Points = userPoint.Points - transactionDomain.Price
				userPoint.CostPoints = userPoint.CostPoints + transactionDomain.Price
				err = s.repo.UpdateUserPoints(c, userPoint)
				if err != nil {
					return nil, err
				}
				product.Stock = product.Stock - 1
				product, err = s.repo.UpdateProduct(c, product.ID, product)
				if err != nil {
					return nil, err
				}
			} else if transaction.PaymentMethod == "buy" {
				// midtrans payment gateway logic
				// if midtrans payment gateway succes then update user points ( increase cost points ), update stock ( decrease stock ), update serial status ( unavailable )
				return nil, errors.New("midtrans payment gateway not implemented yet")
			}

			err = s.repo.UpdateSerialStatus(c, transactionDomain.SerialNumber, "unavailable")
			if err != nil {
				return nil, err
			}
			result, err := s.repo.CreateTransaction(c, transactionDomain)
			if err != nil {
				return nil, err
			}

			return &entity.TransactionsView{
				ID:            result.ID,
				CreatedAt:     result.CreatedAt,
				UpdatedAt:     result.UpdatedAt,
				PaymentMethod: result.PaymentMethod,
				UserID:        result.UserID,
				Username:      userDomain.Username,
				ProductID:     product.ID,
				ProductName:   product.Name,
				Category:      product.Category,
				SerialNumber:  result.SerialNumber,
				IdentifierNum: result.IdentifierNum,
				Price:         result.Price,
				Status:        result.Status,
			}, nil
		} else {
			return nil, errors.New("product payment method not found")
		}
	}
	return nil, err
}

func (s *Service) CreateTransactionByAdmin(c echo.Context, transaction entity.Transactions) (*entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transaction.ID = (guuid.New()).String()
		product, err := s.repo.GetProductByID(c, transaction.ProductID)
		if err != nil {
			return nil, err
		}
		if product.Stock == 0 {
			return nil, errors.New("product out of stock")
		}

		err = helper.ValidateIdentifierNumber(transaction.IdentifierNum)
		if err != nil {
			return nil, err
		}

		if transaction.PaymentMethod == "buy" && product.Buy || transaction.PaymentMethod == "redeem" && product.Redeem {
			var userAuth *entity.Users
			if transaction.UserID != "" {
				userAuth, err = s.repo.GetUserByID(c, transaction.UserID)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.New("user not found")
			}
			if userAuth.Role != "user" {
				return nil, errors.New("user not found")
			}

			userPoint, err := s.repo.GetUserPoints(c, userAuth.ID)
			if err != nil {
				return nil, err
			}

			var result *entity.Transactions
			if transaction.PaymentMethod == "redeem" && userPoint.Points < transaction.Price {
				transaction.Status = "failed"
				return nil, errors.New("user not enough points")
			} else if transaction.PaymentMethod == "redeem" && userPoint.Points >= transaction.Price {
				transaction.Price = product.Price
				serial, err := s.repo.GetSerialNumber(c, transaction.ProductID)
				if err != nil {
					return nil, errors.New("serial number not found")
				}
				transaction.SerialNumber = serial.Serial
				transaction.Status = "success"
				result, err = s.repo.CreateTransaction(c, &transaction)
				if err != nil {
					return nil, err
				}
				err = s.repo.UpdateSerialStatus(c, transaction.SerialNumber, "unavailable")
				if err != nil {
					return nil, err
				}
				product.Stock = product.Stock - 1
				product, err = s.repo.UpdateProduct(c, product.ID, product)
				if err != nil {
					return nil, err
				}

				userPoint.Points = userPoint.Points - transaction.Price
				userPoint.CostPoints = userPoint.CostPoints + transaction.Price
				err = s.repo.UpdateUserPoints(c, userPoint)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.New("payment method not found")
			}

			return &entity.TransactionsView{
				ID:            result.ID,
				CreatedAt:     result.CreatedAt,
				UpdatedAt:     result.UpdatedAt,
				PaymentMethod: result.PaymentMethod,
				UserID:        userAuth.ID,
				Username:      userAuth.Username,
				ProductID:     product.ID,
				ProductName:   product.Name,
				Category:      product.Category,
				SerialNumber:  result.SerialNumber,
				IdentifierNum: result.IdentifierNum,
				Price:         result.Price,
				Status:        result.Status,
			}, nil
		} else {
			return nil, errors.New("product payment method not found")
		}
	}
	return nil, err
}

func (s *Service) UpdateTransactionByAdmin(c echo.Context, ID string, transaction entity.UpdateTransactionBinding) (*entity.TransactionsView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		var (
			validPaymentMethod = false
			validStatus        = false
		)
		if transaction.PaymentMethod == "buy" || transaction.PaymentMethod == "redeem" {
			validPaymentMethod = true
		}
		if transaction.Status == "success" || transaction.Status == "failed" || transaction.Status == "pending" {
			validStatus = true
		}
		err = helper.ValidateIdentifierNumber(transaction.IdentifierNum)
		if err != nil {
			return nil, err
		}
		if !validPaymentMethod {
			return nil, errors.New("payment method not found")
		}
		if !validStatus {
			return nil, errors.New("status not found")
		}

		transactionDomain, err := s.repo.GetTransactionByID(c, ID)
		if err != nil {
			return nil, err
		}
		var userDomain *entity.Users
		var productDomain *entity.Products

		if transaction.PaymentMethod != "" && transactionDomain.PaymentMethod != transaction.PaymentMethod {
			transactionDomain.PaymentMethod = transaction.PaymentMethod
		}
		if transaction.UserID != "" && transactionDomain.UserID != transaction.UserID {
			userDomain, err = s.repo.GetAuth(c, transaction.UserID)
			if err != nil {
				return nil, err
			}
			transactionDomain.UserID = userDomain.ID
		}
		if transaction.ProductID != "" && transactionDomain.ProductID != transaction.ProductID {
			productDomain, err = s.repo.GetProductByID(c, transaction.ProductID)
			if err != nil {
				return nil, err
			}

			if !productDomain.Buy && transactionDomain.PaymentMethod == "buy" || !productDomain.Redeem && transactionDomain.PaymentMethod == "redeem" {
				return nil, errors.New("product payment method not found")
			}

			transactionDomain.ProductID = productDomain.ID
			err = s.repo.UpdateSerialStatus(c, transactionDomain.SerialNumber, "available")
			if err != nil {
				return nil, err
			}
			serial, err := s.repo.GetSerialNumber(c, productDomain.ID)
			if err != nil {
				return nil, err
			}
			transactionDomain.SerialNumber = serial.Serial
			err = s.repo.UpdateSerialStatus(c, transactionDomain.SerialNumber, "unavailable")
			if err != nil {
				return nil, err
			}
			transactionDomain.Price = productDomain.Price
		}

		if transaction.IdentifierNum != "" && transactionDomain.IdentifierNum != transaction.IdentifierNum {
			transactionDomain.IdentifierNum = transaction.IdentifierNum
		}
		if transaction.Status != "" && transactionDomain.Status != transaction.Status {
			if transaction.Status == "success" || transaction.Status == "failed" || transaction.Status == "pending" {
				transactionDomain.Status = transaction.Status
			} else {
				return nil, errors.New("status not found")
			}
		}
		result, err := s.repo.UpdateTransaction(c, ID, transactionDomain)
		if err != nil {
			return nil, err
		}
		return &entity.TransactionsView{
			ID:            result.ID,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
			PaymentMethod: result.PaymentMethod,
			UserID:        userDomain.ID,
			Username:      userDomain.Username,
			ProductID:     productDomain.ID,
			ProductName:   productDomain.Name,
			Category:      productDomain.Category,
			SerialNumber:  result.SerialNumber,
			IdentifierNum: result.IdentifierNum,
			Price:         result.Price,
			Status:        result.Status,
		}, nil
	}
	return nil, err
}

func (s *Service) DeleteTransactionByAdmin(c echo.Context, transactionID string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactionDomain, err := s.repo.GetTransactionByID(c, transactionID)
		if transactionDomain != nil {
			err = s.repo.DeleteTransaction(c, transactionID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}

func (s *Service) GetCountTransactions(c echo.Context) (*entity.GetTransactionsCountView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, _ := s.repo.GetAdminAuth(c, user)
	if adminAuth == nil {
		return nil, errors.New("Unauthorized")
	}

	userCount, err := s.repo.GetCountTransactions(c)

	if err != nil {
		return nil, err
	}
	return userCount, nil
}

func (s *Service) CreateMidtransTransaction(c echo.Context, transaction entity.MidtransTransactionBinding) (*entity.MidtransTransactionView, error) {
	var transactionDomain entity.Transactions

	userDomain, err := s.repo.GetUserByID(c, transaction.UserId)
	if err != nil {
		return nil, err
	}

	productDomain, err := s.repo.GetProductByID(c, transaction.ProductID)
	if err != nil {
		return nil, err
	}
	var snapServer = snap.Client{}
	snapServer.New(os.Getenv("SECRET_KEY"), midtrans.Sandbox)
	custAddress := &midtrans.CustomerAddress{
		FName:       userDomain.Username,
		LName:       userDomain.Username,
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	orderId := (guuid.New()).String()
	callbackUrl := "https://staging-be-loyalty-point-agent.herokuapp.com/api/transactions/payment-success?order_id" + orderId
	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(productDomain.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    userDomain.Username,
			LName:    userDomain.Username,
			Email:    userDomain.Email,
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    productDomain.ID,
				Price: int64(productDomain.Price),
				Qty:   1,
				Name:  productDomain.Name,
			},
		},
		Callbacks: &snap.Callbacks{
			Finish: callbackUrl,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, err := snapServer.CreateTransaction(snapReq)

	transactionDomain.ID = snapResp.Token
	transactionDomain.Status = "pending"
	transactionDomain.UserID = transaction.UserId
	transactionDomain.ProductID = transaction.ProductID
	transactionDomain.SerialNumber = 817238172382
	transactionDomain.IdentifierNum = "085159794050"
	transactionDomain.Price = productDomain.Price
	transactionDomain.PaymentMethod = "buy"

	_, err = s.repo.CreateTransaction(c, &transactionDomain)
	if err != nil {
		return nil, err
	}

	return &entity.MidtransTransactionView{
		Token:     snapResp.Token,
		DirectUrl: snapResp.RedirectURL,
	}, nil
}

func (s *Service) HandlingPaymentStatus(c echo.Context, token string, status int) (string, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	_, err := s.repo.GetAuth(c, user)
	if err != nil {
		return "", err
	}

	transactionDomain, err := s.repo.GetTransactionByID(c, token)

	if status == 409 {
		transactionDomain.Status = "success"
	} else if status == 406 {
		transactionDomain.Status = "pending"
	}

	_, err = s.repo.UpdateTransaction(c, token, transactionDomain)
	if err != nil {
		return "", err
	}
	return "Transaksi Berhasil", nil
}
