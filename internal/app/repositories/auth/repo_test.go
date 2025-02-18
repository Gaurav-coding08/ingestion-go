package auth_test

import (
	"testing"

	"github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/auth"
	repoModels "github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "Failed to connect to test database")

	err = db.AutoMigrate(&repoModels.User{})
	assert.NoError(t, err, "Failed to migrate schema")

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := auth.New(db)

	user := &repoModels.User{
		Email:    "test@example.com",
	}

	createdUser, err := repo.Create(user)

	assert.NoError(t, err, "Expected user to be created successfully")
	assert.NotNil(t, createdUser, "User should not be nil")
	assert.Equal(t, user.Email, createdUser.Email, "User email should match")
}

func TestGetUserByEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := auth.New(db)

	user := &repoModels.User{
		Email:    "test@example.com",
	}

	_, err := repo.Create(user)
	assert.NoError(t, err, "Failed to create test user")

	fetchedUser, err := repo.GetByEmail("test@example.com")

	assert.NoError(t, err, "Expected no error fetching user")
	assert.NotNil(t, fetchedUser, "Fetched user should not be nil")
	assert.Equal(t, user.Email, fetchedUser.Email, "User email should match")
}