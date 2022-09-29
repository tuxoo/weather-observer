package repository

import (
	"context"
	"github.com/tuxoo/weather-observer/internal/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	db *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) *SessionRepository {
	return &SessionRepository{
		db: db.Collection(sessionCollection),
	}
}

func (r *SessionRepository) Save(ctx context.Context, session entity.Session) error {
	_, err := r.db.InsertOne(ctx, session)
	return err
}

func (r *SessionRepository) FindAllByUserId(ctx context.Context, userId string) ([]entity.Session, error) {
	var sessions []entity.Session
	cur, err := r.db.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, err
}

func (r *SessionRepository) DeleteAllByUserId(ctx context.Context, userId string) (int, error) {
	res, err := r.db.DeleteMany(ctx, bson.M{"userId": userId})
	return int(res.DeletedCount), err
}
