package entity

import (
	. "github.com/google/uuid"
	"time"
)

type Session struct {
	Id           string    `bson:"_id,omitempty"`
	RefreshToken UUID      `bson:"refreshToken,omitempty"`
	ExpiresAt    time.Time `bson:"expiresAt,omitempty"`
	UserId       string    `bson:"userId,omitempty"`
}
