package service

import (
	"context"
	"github.com/tuxoo/idler/pkg/hash"
	"time"
	"weather-observer/internal/model/dto"
	"weather-observer/internal/model/entity"
	"weather-observer/internal/repository"
)

type UserService struct {
	repository repository.Users
	hasher     hash.PasswordHasher
}

func NewUserService(repository repository.Users, hasher hash.PasswordHasher) *UserService {
	return &UserService{
		repository: repository,
		hasher:     hasher,
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

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (*dto.LoginResponse, error) {
	return nil, nil
}
