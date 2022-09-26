package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"weather-observer/internal/model/entity"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db.Collection(userCollection),
	}
}

func (r *UserRepository) Save(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}
