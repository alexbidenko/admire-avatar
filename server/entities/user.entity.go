package entities

import (
	"admire-avatar/config"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type BaseUser struct {
	config.Model
	Email string `json:"email"`
	Name  string `json:"name"`
}

type User struct {
	BaseUser
	Password string `json:"password"`
}

func (u *User) GenerateAccessToken() (string, error) {
	var ctx = context.Background()
	expiresAt := time.Now().Add(time.Hour)
	claims := UserClaims{
		Uuid:       uuid.New().String(),
		Authorized: true,
		BaseUser:   u.BaseUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.AccessSecret))
	if err != nil {
		return "", err
	}
	now := time.Now()
	err = config.RedisClient.Set(ctx, claims.Uuid, u.ID, expiresAt.Sub(now)).Err()

	return tokenString, err
}

func (u *User) GenerateRefreshToken() (string, error) {
	var ctx = context.Background()
	expiresAt := time.Now().Add(time.Hour * 24 * 30)
	claims := UserClaims{
		Uuid:     uuid.New().String(),
		BaseUser: u.BaseUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.RefreshSecret))
	if err != nil {
		return "", err
	}
	now := time.Now()
	err = config.RedisClient.Set(ctx, claims.Uuid, u.ID, expiresAt.Sub(now)).Err()

	return tokenString, err
}

type UserClaims struct {
	Authorized bool
	Uuid       string
	BaseUser
	jwt.StandardClaims
}
