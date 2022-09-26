package dto

import (
	. "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}

type SignUpDTO struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=64" example:"alex"`
	LastName  string `json:"lastName" binding:"required,min=2,max=64" example:"cross"`
	Email     string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password  string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}

type User struct {
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	RegisteredAt time.Time `json:"registeredAt"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken UUID   `json:"refreshToken"`
	User         User   `json:"user"`
}
