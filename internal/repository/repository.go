package repository

import (
	"context"
	"github.com/tuxoo/weather-observer/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection    = "users"
	sessionCollection = "sessions"
)

type Users interface {
	Save(ctx context.Context, user entity.User) error
	FindByCredentials(ctx context.Context, email, passwordHash string) (*entity.User, error)
	FindById(ctx context.Context, id string) (*entity.User, error)
}

type Sessions interface {
	Save(ctx context.Context, session entity.Session) error
	FindAllByUserId(ctx context.Context, userId string) ([]entity.Session, error)
	DeleteAllByUserId(ctx context.Context, userId string) (int, error)
}

type Repositories struct {
	Users    Users
	Sessions Sessions
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:    NewUserRepository(db),
		Sessions: NewSessionRepository(db),
	}
}
