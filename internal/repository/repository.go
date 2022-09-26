package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"weather-observer/internal/model/entity"
)

const (
	userCollection = "users"
)

type Users interface {
	Save(ctx context.Context, user entity.User) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
