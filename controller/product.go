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
