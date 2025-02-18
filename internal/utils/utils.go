package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Email  string    `json:"email"`
	UserID uuid.UUID `json:"user_id"`
	Type   string    `json:"type"`
	jwt.StandardClaims
}

func GenerateJWT(
	email string,
	userID uuid.UUID,
	expirationTime time.Duration,
	tokenType string,
) (string, error) {
	now := time.Now().UTC()
	claims := &Claims{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(expirationTime).Unix(),
			IssuedAt:  now.Unix(),
		},
		Type: tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(
	tokenString string,
	expectedType string,
) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Type != expectedType {
		return nil, fmt.Errorf("invalid token type: expected %s, got %s", expectedType, claims.Type)
	}

	if err != nil {
		return nil, parseValidationError(err)
	}

	return claims, nil
}

// parseValidationError processes JWT validation errors
func parseValidationError(err error) error {
	if validationErr, ok := err.(*jwt.ValidationError); ok {
		switch {
		case validationErr.Errors&jwt.ValidationErrorExpired != 0:
			return errors.New("token is expired")
		case validationErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
			return errors.New("invalid token signature")
		}
	}
	return fmt.Errorf("error parsing token: %v", err)
}
