package entity

import "time"

type Role string

const (
	USER_ROLE  = Role("USER")
	ADMIN_ROLE = Role("ADMIN")
)

type User struct {
	Id           string    `bson:"_id,omitempty"`
	FirstName    string    `bson:"firstName,omitempty"`
	LastName     string    `bson:"lastName,omitempty"`
	Email        string    `bson:"email,omitempty"`
	PasswordHash string    `bson:"passwordHash,omitempty"`
	RegisteredAt time.Time `bson:"registeredAt,omitempty"`
	VisitedAt    time.Time `bson:"visitedAt,omitempty"`
	Role         Role      `bson:"role,omitempty"`
	IsEnable     bool      `bson:"isEnable,omitempty"`
}
