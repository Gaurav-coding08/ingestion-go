// Auth Middleware for authenticating an admin user which is only allowed to send updates.
// Will be imported in many services.

package v1

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Email  string    `json:"email"`
	UserID uuid.UUID `json:"user_id"`
	Type   string    `json:"type"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		// as only "access_tokens" can be validated not refresh token, so if in user-auth-go
		// the str value of AccessToken changes, then here also it should be changed.
		claims, err := ValidateToken(tokenString, "access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
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
