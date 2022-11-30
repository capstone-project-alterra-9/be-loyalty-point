package repository

import (
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetTransactions(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsRedeem(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "payment_method = ?", "redeem").Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsBuy(c echo.Context) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "payment_method = ?", "buy").Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetHistory(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetHistoryBuy(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ? AND payment_method = ?", ID, "buy").Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetHistoryRedeem(c echo.Context, ID string) ([]entity.Transactions, error) {
	var transactions []entity.Transactions
	err := r.connection.Find(&transactions, "user_id = ? AND payment_method = ?", ID, "redeem").Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) CreateTransaction(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error) {
	err := r.connection.Create(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *repository) GetTransactionByID(c echo.Context, ID string) (*entity.Transactions, error) {
	var transactionDomain *entity.Transactions
	err := r.connection.First(&transactionDomain, "id = ?", ID).Error
	if transactionDomain.ID == "" {
		return nil, err
	}
	return transactionDomain, nil
}

func (r *repository) UpdateTransaction(c echo.Context, transaction *entity.Transactions) (*entity.Transactions, error) {
	err := r.connection.Save(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *repository) DeleteTransaction(c echo.Context, ID string) error {
	var transactionDomain *entity.Transactions
	err := r.connection.First(&transactionDomain, "id = ?", ID).Error
	if transactionDomain.ID == "" {
		return err
	}
	err = r.connection.Delete(transactionDomain).Error
	if err != nil {
		return err
	}
	return nil
}
