package service

import (
	"capstone-project/entity"
	"capstone-project/helper"
	jwtAuth "capstone-project/middleware"
	"errors"
	"math/rand"
	"time"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// same data product gaboleh -- uda
// empty name gaboleh -- uda
// descripsi bole kosong kata era -- uda
// stock 0 boleh kata era -- uda
// stock negatif gaboleh -- uda
// harga negatif gaboleh -- uda
// update pake data sama gaboleh -- uda
// update product tanpa nama gaboleh -- uda

func (s *Service) CreateProduct(c echo.Context, product *entity.Products) (*entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		if product.Name == "" || product.Category == "" || product.Price == 0 {
			return nil, helper.ErrEmptyData
		}

		product.ID = guuid.New().String()
		product.Redeem = true
		if product.Category == "credits" || product.Category == "data-quota" {
			product.Buy = true
		} else if product.Category == "e-money" || product.Category == "cashout" {
			product.Buy = false
		} else {
			return nil, errors.New("invalid category")
		}

		if product.Price < 0 {
			return nil, errors.New("invalid price")
		}
		if product.Stock < 0 {
			return nil, errors.New("invalid stock")
		}

		rand.Seed(time.Now().UnixNano())
		for i := 0; i < product.Stock; i++ {
			randomNum := rand.Int63n(999999999999-99999999999) - 99999999999
			if randomNum < 0 {
				randomNum = randomNum * -1
			}
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

func (s *Service) GetProductsByMethod(c echo.Context, method string) ([]entity.Products, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		if method == "buy" || method == "redeem" {
			return s.repo.GetProductsByMethod(c, method)
		} else {
			return nil, errors.New("invalid method")
		}
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
		productDomain, err := s.repo.GetProductByID(c, ID)
		if productDomain != nil {
			if product.Name == productDomain.Name && product.Description == productDomain.Description && product.Price == productDomain.Price && product.Category == productDomain.Category && product.Stock == productDomain.Stock && product.Image == productDomain.Image {
				return nil, helper.ErrSameDataRequest
			}
			if product.Name == "" {
				return nil, helper.ErrEmptyData
			}
			if product.Price < 0 {
				return nil, errors.New("invalid price")
			}
			if product.Stock < 0 {
				return nil, errors.New("invalid stock")
			}
			var diffBuyMethod bool
			if product.Category != productDomain.Category {
				// product.Redeem = true
				if product.Category == "credits" || product.Category == "data-quota" {
					product.Buy = true
				} else if product.Category == "e-money" || product.Category == "cashout" {
					product.Buy = false
				} else {
					return nil, errors.New("invalid category")
				}
			}
			if product.Buy != productDomain.Buy {
				diffBuyMethod = true
			}
			if product.Stock != productDomain.Stock {
				if product.Stock > productDomain.Stock {
					rand.Seed(time.Now().UnixNano())
					for i := 0; i < product.Stock-productDomain.Stock; i++ {
						randomNum := rand.Int63n(999999999999-99999999999) - 99999999999
						if randomNum < 0 {
							randomNum = randomNum * -1
						}
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
				} else if product.Stock < productDomain.Stock {
					diff := productDomain.Stock - product.Stock
					err := s.repo.DeleteNSerialNumberByProductID(c, ID, diff)
					if err != nil {
						return nil, err
					}
				}
			}

			if diffBuyMethod { // this function created because if true value updates with true value, it will be false value
				err := s.repo.UpdateMethodProduct(c, ID, product.Buy)
				if err != nil {
					return nil, err
				}
			}
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
			err := s.repo.DeleteAllSerialNumberByProductID(c, ID)
			if err != nil {
				return err
			}
			err = s.repo.DeleteProduct(c, ID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}

func (s *Service) GetCountProducts(c echo.Context) (*entity.GetProductsCountView, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth == nil {
		return nil, errors.New("Unauthorized")
	}

	userCount, err := s.repo.GetCountProducts(c)

	if err != nil {
		return nil, err
	}
	return userCount, nil;
}
