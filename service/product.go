package service

import (
	"capstone-project/entity"
	"capstone-project/middleware"
	"errors"
	"math/rand"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	result, err := s.repo.GetAdminAuth(c)
	if err != nil {
		return nil, err
	}
	for _, admin := range result {
		if admin.Username == middleware.ExtractTokenUsername(c) {
			product.CreatedBy = middleware.ExtractTokenUsername(c)
			product.UpdatedBy = product.CreatedBy
			SerialList, err := s.repo.GetSerialNumbers(c)
			if err != nil {
				return nil, err
			}
			for i := 0; i < product.Stock; i++ {
				var randomNum string
				for {
					randomNum = strconv.Itoa(rand.Intn(999999999999-99999999999) - 99999999999)
					for _, serial := range SerialList {
						if randomNum == serial.Serial {
							break
						}
					}
					break
				}

				serialNumber := &entity.SerialNumbers{
					ProductID: product.ID,
					Serial:    randomNum,
					Status:    "Available",
					CreatedBy: product.CreatedBy,
				}
				err := s.repo.CreateSerialNumber(c, serialNumber)
				if err != nil {
					return nil, err
				}
				SerialList = append(SerialList, *serialNumber)
			}
			return s.repo.CreateProduct(c, product)
		}
	}

	return nil, errors.New("you are not an admin")
}
