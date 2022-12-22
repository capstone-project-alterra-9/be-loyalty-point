package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	result, err := Service.GetTransactions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetTransactionsByMethod(c echo.Context) error {
	method := c.Param("paymentMethod")
	result, err := Service.GetTransactionsByMethod(c, method)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	result, err := Service.GetTransactionByID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetHistory(c echo.Context) error {
	result, err := Service.GetHistory(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetHistoryByMethod(c echo.Context) error {
	method := c.Param("paymentMethod")
	result, err := Service.GetHistoryByMethod(c, method)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetHistoryByMethodCategory(c echo.Context) error {
	method := c.Param("paymentMethod")
	category := c.Param("category")
	result, err := Service.GetHistoryByMethodCategory(c, method, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func CreateTransactionByUser(c echo.Context) error {
	var transaction entity.TransactionsBinding
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.CreateTransactionByUser(c, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func CreateTransactionByAdmin(c echo.Context) error {
	var transaction entity.Transactions
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.CreateTransactionByAdmin(c, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func UpdateTransactionByAdmin(c echo.Context) error {
	id := c.Param("id")
	var transaction entity.UpdateTransactionBinding
	err := c.Bind(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.UpdateTransactionByAdmin(c, id, transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func DeleteTransactionByAdmin(c echo.Context) error {
	id := c.Param("id")
	err := Service.DeleteTransactionByAdmin(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
}

func GetCountTransactions(c echo.Context) error {
	result, err := Service.GetCountTransactions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func CreateMidtransTransaction(c echo.Context) error {
	var midtransTransaction entity.MidtransTransactionBinding
	if err := c.Bind(&midtransTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.CreateMidtransTransaction(c, midtransTransaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}
