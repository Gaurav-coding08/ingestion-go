package v1

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
}

type AuthTokenResponse struct {
	AccessToken string        `json:"access_token"`
	Type        string        `json:"type"`
	ExpiresAt   time.Duration `json:"expires_at"`
}
