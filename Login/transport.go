package login

import (
	"context"
	"encoding/json"
	"errors"
	auth "lms/Auth"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/* --------------------------- Request & Response --------------------------- */
type SignupRequest struct {
	TypeId int `json:"typeId"`
	user   Users
}
type SignupResponse struct {
	V   string `json:"Response:,omitempty"`
	Err string `json:"error,omitempty"`
}
type LoginRequest struct {
	credentials Credentials
}
type LoginResponse struct {
	loggedUser LoggedInUser
	err        string
}

/* ----------------------------- Decode Request ----------------------------- */
func DecodeSignupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SignupRequest
	params := mux.Vars(r)
	switch params["type"] {
	case "employee":
		request.TypeId = 3
	case "manager":
		request.TypeId = 2
	case "admin":
		request.TypeId = 1
	default:
		return nil, errors.New("Invalid Type")
	}
	if err := json.NewDecoder(r.Body).Decode(&request.user); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request.credentials); err != nil {
		return nil, err
	}
	return request, nil
}

// /Encode Response contains jwt token string
func EncodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	res := response.(LoginResponse)
	if res.err != "" {
		return json.NewEncoder(w).Encode(res.err)
	}
	CookieName := "logintoken"
	err := auth.SetCookie(w, CookieName, res.loggedUser.Token, time.Hour*24, "/", true, true)
	if err != nil {
		http.Error(w, "Failed to set cookie", http.StatusInternalServerError)
		return nil
	}


	return json.NewEncoder(w).Encode(res.loggedUser.Details)
}
