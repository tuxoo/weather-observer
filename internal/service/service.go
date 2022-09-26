package service

import (
	"context"
	"github.com/tuxoo/idler/pkg/hash"
	"weather-observer/internal/model/dto"
	"weather-observer/internal/repository"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (*dto.LoginResponse, error)
	//GetById(ctx context.Context, id UUID) (*dto.UserDTO, error)
	//GetAll(ctx context.Context) ([]dto.UserDTO, error)
	//GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Services struct {
	UserService Users
}

func NewServices(repository *repository.Repositories, hasher hash.PasswordHasher) *Services {
	userService := NewUserService(repository.Users, hasher)

	return &Services{
		UserService: userService,
	}
}
