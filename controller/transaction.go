package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	transactions, err := Service.GetTransactions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transactions", transactions))
}

func GetTransactionsRedeem(c echo.Context) error {
	transactions, err := Service.GetTransactionsRedeem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Succes get transactions by method redeem", transactions))
}

func GetTransactionsBuy(c echo.Context) error {
	transactions, err := Service.GetTransactionsBuy(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transactions", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Succes get transactions by method buy", transactions))
}

func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	result, err := Service.GetTransactionByID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get transaction", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get transaction", result))
}

func GetHistory(c echo.Context) error {
	transactions, err := Service.GetHistory(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get history", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get history", transactions))
}

func GetHistoryBuy(c echo.Context) error {
	transactions, err := Service.GetHistoryBuy(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get buy history", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get buy history", transactions))
}

func GetHistoryRedeem(c echo.Context) error {
	transactions, err := Service.GetHistoryRedeem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get redeem history", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get redeem history", transactions))
}

func CreateTransactionByUser(c echo.Context) error {
	var transaction entity.TransactionsBinding
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.CreateTransactionByUser(c, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to add transaction", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success add transaction", result))
}

func CreateTransactionByAdmin(c echo.Context) error {
	var transaction entity.Transactions
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.CreateTransactionByAdmin(c, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to add transaction", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success add transaction", result))
}

func UpdateTransactionByAdmin(c echo.Context) error {
	id := c.Param("id")
	var transaction entity.UpdateTransactionBinding
	err := c.Bind(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.UpdateTransactionByAdmin(c, id, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to update transaction", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success update transaction", result))
}

func DeleteTransactionByAdmin(c echo.Context) error {
	id := c.Param("id")
	err := Service.DeleteTransactionByAdmin(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to delete transaction", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success delete transaction", dto.EmptyObj{}))
}
