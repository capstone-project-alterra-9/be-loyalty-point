package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"
	"errors"
	"math/rand"
	"time"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Service) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		product.ID = guuid.New().String()
		product.Redeem = true
		if product.Category == "credits" || product.Category == "data-quota" {
			product.Buy = true
		} else if product.Category == "e-money" || product.Category == "cashout" {
			product.Buy = false
		} else {
			return nil, errors.New("invalid category")
		}
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

func (s *Service) GetProducts(c echo.Context) ([]entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		products, err := s.repo.GetProducts(c)
		if err != nil {
			return nil, err
		}
		return products, nil
	}
	return nil, err
}

func (s *Service) GetProductsByCategory(c echo.Context, category string) ([]entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		if category == "credits" || category == "data-quota" || category == "e-money" || category == "cashout" {
			return s.repo.GetProductsByCategory(c, category)
		} else {
			return nil, errors.New("invalid category")
		}
	}
	return nil, err
}

func (s *Service) GetProductByID(c echo.Context, ID string) (*entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		product, err := s.repo.GetProductByID(c, ID)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
	return nil, err
}

func (s *Service) UpdateProduct(c echo.Context, ID string, product *entity.Products) (*entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		updateProduct, err := s.repo.GetProductByID(c, ID)
		if updateProduct != nil {
			if product.Category == "credits" || product.Category == "data-quota" {
				product.Buy = true
			} else if product.Category == "e-money" || product.Category == "cashout" {
				product.Buy = false
			} else {
				return nil, errors.New("invalid category")
			}
			product.ID = ID
			return s.repo.UpdateProduct(c, ID, product)
		}
		return nil, err
	}
	return nil, err
}

func (s *Service) DeleteProduct(c echo.Context, ID string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		product, err := s.repo.GetProductByID(c, ID)
		if product != nil {
			err := s.repo.DeleteProduct(c, ID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}
