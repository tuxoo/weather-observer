package service

import (
	"context"
	. "github.com/google/uuid"
	"github.com/tuxoo/weather-observer/internal/config"
	"github.com/tuxoo/weather-observer/internal/model/entity"
	"github.com/tuxoo/weather-observer/internal/repository"
	"time"
)

type SessionService struct {
	repository repository.Sessions
	cfg        *config.Config
}

func NewSessionService(repository repository.Sessions, cfg *config.Config) *SessionService {
	return &SessionService{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *SessionService) Create(ctx context.Context, userId string) (refreshToken UUID, err error) {
	refreshToken = New()
	sessions, err := s.getAllByUserId(ctx, userId)
	if err != nil {
		return refreshToken, err
	}

	sessionsNum := len(sessions)

	if sessionsNum >= s.cfg.Auth.SessionMax {
		if count, err := s.deleteAll(ctx, userId); err != nil && count != sessionsNum {
			return refreshToken, err
		}
	}

	session := entity.Session{
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(s.cfg.Auth.RefreshTokenTTL),
		UserId:       userId,
	}

	if err := s.repository.Save(ctx, session); err != nil {
		return refreshToken, nil
	}

	return refreshToken, nil
}

func (s *SessionService) getAllByUserId(ctx context.Context, userId string) ([]entity.Session, error) {
	return s.repository.FindAllByUserId(ctx, userId)
}

func (s *SessionService) deleteAll(ctx context.Context, userId string) (int, error) {
	return s.repository.DeleteAllByUserId(ctx, userId)
}
