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
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.CreateProduct(c, &product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to create product", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success create product", result))
}

func GetProducts(c echo.Context) error {
	result, err := Service.GetProducts(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to get products", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get products", result))
}

func GetProductsByCategory(c echo.Context) error {
	category := c.Param("categoryName")
	result, err := Service.GetProductsByCategory(c, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to get products", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get products", result))
}

func GetProductByID(c echo.Context) error {
	id := c.Param("id")
	result, err := Service.GetProductByID(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to get product", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success get product", result))
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product entity.Products
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to bind request", err))
	}
	result, err := Service.UpdateProduct(c, id, &product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to update product", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success update product", result))
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := Service.DeleteProduct(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse("Failed to delete product", err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse("Success delete product", nil))
}

// func GetCountProducts(c echo.Context) error {
// 	result, err := Service.GetCountUsers(c)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse("Failed to get total users", err))
// 	}
// 	return c.JSON(http.StatusOK, dto.BuildResponse("Success get total users", result))
// }