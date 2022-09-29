package service

import (
	"context"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/idler/pkg/cache"
	"github.com/tuxoo/idler/pkg/hash"
	"github.com/tuxoo/weather-observer/internal/config"
	"github.com/tuxoo/weather-observer/internal/model/dto"
	"github.com/tuxoo/weather-observer/internal/model/entity"
	"github.com/tuxoo/weather-observer/internal/repository"
	"time"
)

type UserService struct {
	repository     repository.Users
	cfg            *config.Config
	hasher         hash.PasswordHasher
	tokenManager   auth.TokenManager
	sessionService Sessions
	cache          cache.Cache[string, entity.User]
}

func NewUserService(repository repository.Users, cfg *config.Config, hasher hash.PasswordHasher, tokenManager auth.TokenManager, sessionService Sessions, cache cache.Cache[string, entity.User]) *UserService {
	return &UserService{
		repository:     repository,
		cfg:            cfg,
		hasher:         hasher,
		tokenManager:   tokenManager,
		sessionService: sessionService,
		cache:          cache,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto dto.SignUpDTO) error {
	user := entity.User{
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Email:        dto.Email,
		PasswordHash: s.hasher.Hash(dto.Password),
		RegisteredAt: time.Now(),
		VisitedAt:    time.Now(),
		Role:         entity.USER_ROLE,
		IsEnable:     false,
	}

	return s.repository.Save(ctx, user)
}

func (s *UserService) SignIn(ctx context.Context, inDTO dto.SignInDTO) (response dto.LoginResponse, err error) {
	user, err := s.repository.FindByCredentials(ctx, inDTO.Email, s.hasher.Hash(inDTO.Password))
	if err != nil {
		return response, err
	}

	refreshToken, err := s.sessionService.Create(ctx, user.Id)
	if err != nil {
		return response, err
	}

	accessToken, err := s.tokenManager.GenerateToken(user.Id, s.cfg.Auth.AccessTokenTTL)

	if err := s.cache.Set(ctx, user.Id, user); err != nil {
		return response, err
	}

	response = dto.LoginResponse{
		AccessToken:  string(accessToken),
		RefreshToken: refreshToken,
		User: dto.User{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			RegisteredAt: user.RegisteredAt,
		},
	}

	return
}

func (s *UserService) GetById(ctx context.Context, id string) (user *entity.User, err error) {
	user, err = s.cache.Get(ctx, id)
	if user == nil && err != nil {
		user, err = s.repository.FindById(ctx, id)
	}
	return
}
