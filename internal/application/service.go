package application

import "github.com/allenliao0119/go-auth-service/internal/service/password"

type Service struct {
	Password *password.Service
}

func NewService() *Service {
	return &Service{
		Password: password.NewService(),
	}
}