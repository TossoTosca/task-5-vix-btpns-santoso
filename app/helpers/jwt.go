package helpers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret_key")

// GenerateToken generates a new JWT token
func GenerateToken(userID uint) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		// Return the secret key
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Get the user ID from the claims
	userID, ok := claims["user_id"].(float64)

	if !ok {
		return 0, fmt.Errorf("invalid user ID in token")
	}

	return uint(userID), nil
}
