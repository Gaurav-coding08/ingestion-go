package auth

import (
	models "github.com/Gaurav-coding08/ingestion-go/internal/app/models"
	repoModels "github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/models"
	v1 "github.com/Gaurav-coding08/ingestion-go/pkg/client"
	"github.com/google/uuid"
	"github.com/Gaurav-coding08/ingestion-go/internal/utils"
)

type Repo interface {
	Create(user *repoModels.User) (*repoModels.User, error)
	GetByEmail(email string) (*repoModels.User, error)
}

type service struct {
	repo Repo
}

func New(repo Repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(
	userRegisterReq v1.CreateUserRequest,
) (models.User, error) {
	user, err := s.repo.Create(&repoModels.User{
		UserID: uuid.New(),
		Name:   userRegisterReq.Name,
		Email:  userRegisterReq.Email,
	})

	if err != nil {
		return models.User{}, err
	}

	return models.User{}.FromRepo(user), nil
}

func (s *service) Login(loginUser models.LoginUser) (*models.AuthToken, error) {
	user, err := s.repo.GetByEmail(loginUser.Email)
	if err != nil {
		return nil, err
	}


	// Generate Access Token
	accessToken, err := utils.GenerateJWT(user.Email, user.UserID, utils.AccessTokenExpiry, utils.AccessToken)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	// Generate Refresh Token.

	// update db with new token

	return &models.AuthToken{
		AccessToken:  accessToken,
		Type:         "Bearer",
		ExpiresAt:    utils.AccessTokenExpiry,
	}, nil
}
