package helpers

import (
	"net/http"
	"strings"
)

// GetAuthorizationToken extracts the JWT token from the Authorization header
func GetAuthorizationToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token
}
