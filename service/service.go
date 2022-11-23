package service

import (
	"capstone-project/repository"
)

type Service struct {
	repo repository.Repo
}

func NewService(repository repository.Repo) Svc {
	return &Service{
		repo: repository,
	}
}

type Svc interface {
	authSvc
	productSvc
	trasanctionSvc
	userSvc
}

type authSvc interface{}

type productSvc interface{}

type trasanctionSvc interface{}

type userSvc interface{}
