package service

import (
	"context"
	"github.com/tuxoo/idler/pkg/auth"
	"weather-observer/internal/model/dto"
	"weather-observer/internal/repository"
)

type UserService struct {
	repository repository.Users
}

func NewUserService(repository repository.Users) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto dto.SignUpDTO) error {
	return nil
}

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (auth.Token, error) {
	return "", nil
}
