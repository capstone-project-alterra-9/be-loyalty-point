package controller

import (
	"net/http"

	"capstone-project/dto"
	"capstone-project/entity"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product entity.Products
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.CreateProduct(c, &product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusCreated, http.StatusText(http.StatusCreated), result))
}

func GetProducts(c echo.Context) error {
	result, err := Service.GetProducts(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetProductsByMethod(c echo.Context) error {
	paymentMethod := c.Param("paymentMethod")
	result, err := Service.GetProductsByMethod(c, paymentMethod)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetProductsByCategory(c echo.Context) error {
	category := c.Param("categoryName")
	result, err := Service.GetProductsByCategory(c, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func GetProductByID(c echo.Context) error {
	id := c.Param("id")
	result, err := Service.GetProductByID(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product entity.Products
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	result, err := Service.UpdateProduct(c, id, &product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := Service.DeleteProduct(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), nil))
}

// func GetCountProducts(c echo.Context) error {
// 	result, err := Service.GetCountUsers(c)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get total users", err))
// 	}
// 	return c.JSON(http.StatusOK, dto.BuildResponse("Success get total users", result))
// }