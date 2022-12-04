package service

import (
	"capstone-project/entity"
	jwtAuth "capstone-project/middleware"

	"errors"
	"github.com/labstack/echo/v4"
)

func (s *Service) DeleteOneById(c echo.Context, transactionID string) error {
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		transactionDomain, err := s.repo.GetTransactionByID(c, transactionID)
		if transactionDomain != nil {
			err = s.repo.DeleteOneById(c, transactionID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}
