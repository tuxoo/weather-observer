package repository

import (
	"context"
	"github.com/tuxoo/weather-observer/internal/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	userDb := db.Collection(userCollection)
	_, err := userDb.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		panic("Error had been happened during creating index")
	}

	return &UserRepository{
		db: userDb,
	}
}

func (r *UserRepository) Save(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, passwordHash string) (user *entity.User, err error) {
	err = r.db.FindOne(ctx, bson.M{
		"email":        email,
		"passwordHash": passwordHash,
	}).Decode(&user)
	return
}

func (r *UserRepository) FindById(ctx context.Context, id string) (user *entity.User, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.db.FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&user)
	return
}
