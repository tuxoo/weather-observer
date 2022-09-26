package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db.Collection(userCollection),
	}
}

func (r *UserRepository) Save(ctx context.Context) {

}
