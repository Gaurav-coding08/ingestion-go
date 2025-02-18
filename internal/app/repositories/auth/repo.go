package auth

import (
	repoModels "github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *repoModels.User) (*repoModels.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetByEmail(email string) (*repoModels.User, error) {
	var user repoModels.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
