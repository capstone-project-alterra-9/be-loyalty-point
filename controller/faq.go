package controller

import (
	"capstone-project/dto"
	"capstone-project/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateFaq(c echo.Context) error {
	var faq entity.FAQPayloadBinding
	if err := c.Bind(&faq); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	result, err := Service.CreateFaq(c, faq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BuildErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err))
	}

	return c.JSON(http.StatusCreated, dto.BuildResponse(http.StatusCreated, http.StatusText(http.StatusCreated), result))
}

func GetFaqs(c echo.Context) error {
	users, err := Service.GetFaqs(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BuildErrorResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, dto.BuildResponse(http.StatusOK, http.StatusText(http.StatusOK), users))
}