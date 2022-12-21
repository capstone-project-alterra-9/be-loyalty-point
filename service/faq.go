package service

import (
	"capstone-project/entity"
	"capstone-project/helper"
	jwtAuth "capstone-project/middleware"

	guuid "github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func (s *Service) CreateFaq(c echo.Context, faq entity.FAQPayloadBinding) (*entity.FAQ, error) {
	var faqDomain entity.FAQ
	user := jwtAuth.ExtractTokenUsername(c)
	adminAuth, err := s.repo.GetAdminAuth(c, user)
	if adminAuth != nil {
		if faq.Title == "" || faq.Category == "" || faq.Body == "" {
			return nil, helper.ErrEmptyData
		}

		faqDomain.ID = guuid.New().String()
		faqDomain.Title = faq.Title
		faqDomain.Category = faq.Category
		faqDomain.Body = faq.Body

		return s.repo.CreateFaq(c, faqDomain)
	}
	return nil, err
}

func (s *Service) GetFaqs(c echo.Context) ([]entity.FAQ, error) {
	user := jwtAuth.ExtractTokenUsername(c)
	auth, err := s.repo.GetAuth(c, user)
	if auth != nil {
		faqs, err := s.repo.GetFaqs(c)
		if err != nil {
			return nil, err
		}
		return faqs, nil
	}
	return nil, err
}