package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"
	"math/rand"
	"time"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Service) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < product.Stock; i++ {
			randomNum := rand.Int63n(999999999999-99999999999) - 99999999999
			serialNumber := &entity.SerialNumbers{
				ID:        guuid.New().String(),
				ProductID: product.ID,
				Serial:    randomNum,
				Status:    "available",
			}
			err := s.repo.CreateSerialNumber(c, serialNumber)
			if err != nil {
				return nil, err
			}
		}
		return s.repo.CreateProduct(c, product)
	}
	return nil, err
}
