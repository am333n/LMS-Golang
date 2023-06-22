package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// EncodeError encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case ErrUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrLoginInvalid:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrLoginInvalid:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrTokenExpired:
		w.WriteHeader(http.StatusForbidden)
	case ErrRefreshTokenExpired:
		w.WriteHeader(http.StatusNotAcceptable)
	case ErrUserNameExist:
		w.WriteHeader(http.StatusBadRequest)
	case ErrIDNotFound:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidCredentials:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var (
	authApi string = ""
)
var (
	// ErrRepo Repo error
	ErrRepo = errors.New("Unable to handle Repo Request")
	// ErrIDNotFound id not found error
	ErrIDNotFound           = errors.New("Id Not found")
	ErrInvalidArgument      = errors.New("Invalid Argument")
	ErrLoginInvalid         = errors.New("Invalid Credentials")
	ErrLoginInactiveUser    = errors.New("Credentials are Inactive")
	ErrInternalServer       = errors.New("Internal Server Error")
	ErrUnauthorized         = errors.New("Unauthorized Access")
	ErrBadRequest           = errors.New("Bad Request")
	ErrTokenExpired         = errors.New("Token Expired")
	ErrRefreshTokenExpired  = errors.New("Refresh Token Expired")
	ErrCookieNotFound       = errors.New("Cookie Not found")
	ErrTokenCorrupted       = errors.New("Invalid/Malformed Auth Token")
	ErrUserNameExist        = errors.New("Username already exists")
	ErrEmailExist           = errors.New("Email already exists")
	ErrRequiredFieldMissing = errors.New("Required Field Missing")
	ErrBadParameter         = errors.New("Bad Paramter Request")
	ErrNoPermission         = errors.New("No Permission")
	ErrNoFile               = errors.New("no file exists")
	ErrInvalidCredentials   = errors.New("Invalid Credentials")
)
