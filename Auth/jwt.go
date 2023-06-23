package auth

import (
	"fmt"
	"lms/common"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("heisenberg")

// GenerateJWTToken generates a new JWT token for the given username.
func GenerateJWTToken(username string, id int, userType int) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() // Set the token expiration time
	claims["userType"] = userType

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return tokenString, nil
}
func SetCookie(w http.ResponseWriter, name, value string, expires time.Duration, path string, httpOnly, secure bool) error {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(expires),
		Path:     path,
		HttpOnly: httpOnly,
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	return nil
}
func extractTokenFromRequest(r *http.Request) (string, error) {
	// Implement the logic to extract the JWT token from the request, e.g., from headers or cookies
	cookie, err := r.Cookie("logintoken")
	if err != nil {
		return "", common.ErrCookieNotFound
	}
	// Return the token value
	return cookie.Value, nil
}
