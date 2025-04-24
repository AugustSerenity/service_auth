package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshToken struct {
	ID          uint
	UserID      string
	AccessJTI   string
	HashedToken string
	IP          string
	Used        bool
	CreatedAt   time.Time
}

type Claims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.RegisteredClaims
}
