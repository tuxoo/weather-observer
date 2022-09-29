package service

import (
	"context"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/idler/pkg/hash"
	"github.com/tuxoo/weather-observer/internal/config"
	"github.com/tuxoo/weather-observer/internal/model/dto"
	"github.com/tuxoo/weather-observer/internal/repository"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (dto.LoginResponse, error)
	//GetById(ctx context.Context, id UUID) (*dto.UserDTO, error)
	//GetAll(ctx context.Context) ([]dto.UserDTO, error)
	//GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Sessions interface {
	Create(ctx context.Context, userId string) (refreshToken UUID, err error)
}

type Services struct {
	UserService    Users
	SessionService Sessions
}

func NewServices(repository *repository.Repositories, hasher hash.PasswordHasher, tokenManager auth.TokenManager, cfg *config.Config) *Services {
	sessionService := NewSessionService(repository.Sessions, cfg)
	return &Services{
		UserService:    NewUserService(repository.Users, cfg, hasher, tokenManager, sessionService),
		SessionService: sessionService,
	}
}
