package models

import (
	"time"

	repoModels "github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/models"
	v1 "github.com/Gaurav-coding08/ingestion-go/pkg/client"
	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID
	Name   string
	Email  string
	Age    *int
}

type AuthToken struct {
	AccessToken  string
	RefreshToken string
	Type         string
	ExpiresAt    time.Duration
}

type LoginUser struct {
	Email string
}

func (u LoginUser) FromRequest(req v1.LoginUserRequest) LoginUser {
	return LoginUser{
		Email: req.Email,
	}
}

func (u User) FromRepo(repoUser *repoModels.User) User {
	return User{
		UserID: repoUser.UserID,
		Name:   repoUser.Name,
		Email:  repoUser.Email,
	}
}

func (at AuthToken) ToResponse() v1.AuthTokenResponse {
	return v1.AuthTokenResponse{
		AccessToken: at.AccessToken,
		Type:        at.Type,
		ExpiresAt:   at.ExpiresAt,
	}
}

func (user User) ToResponse() v1.User {
	return v1.User{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}
}
