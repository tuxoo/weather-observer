package service

import (
	"context"
	"github.com/tuxoo/idler/pkg/auth"
	"weather-observer/internal/model/dto"
	"weather-observer/internal/repository"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	//GetById(ctx context.Context, id UUID) (*dto.UserDTO, error)
	//GetAll(ctx context.Context) ([]dto.UserDTO, error)
	//GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Services struct {
	UserService Users
}

func NewServices(repository *repository.Repositories) *Services {
	userService := NewUserService(repository.Users)

	return &Services{
		UserService: userService,
	}
}
