package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection = "users"
)

type Users interface {
	Save(ctx context.Context)
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
