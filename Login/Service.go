package login

import (
	"encoding/json"
	"errors"
	"fmt"
	dc "lms/Database"
	auth "lms/Auth"
	"lms/common"
	"net/http"
	"time"

	"gorm.io/gorm"
)

/* --------------------------------- Structs -------------------------------- */
type RepoService struct{}
type Users struct {
	gorm.Model
	loginId  int    `gorm:"autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoggedInUser struct {
	Token   string
	Details UserDetails
}
type UserDetails struct {
	Username string
	UserId   int
	UserType int
}

/* -------------------------------- Services -------------------------------- */
type LoginService interface {
	Login(credentials Credentials) (LoggedInUser, error)
	Signup(TypeId int, users Users) (string, error)
}

/* ----------------------------- Singup Function ---------------------------- */

func (RepoService) Signup(TypeId int, users Users) (string, error) {
	switch TypeId {
	case 1:
		users.Type = "Employee"
	case 2:
		users.Type = "Manager"
	case 3:
		users.Type = "Admin"
	}
	db, err := dc.GetDB()
	if err != nil {
		return "", errors.New("dB Connection Failed")
	}
	if err := db.Model(&users).Where("username=?", users.Username).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&users).Error; err != nil {

				return "", errors.New("Could Not Signup")
			}
			return "Signup Successfull", nil
		}
		return "", err
	}
	return " Username Already Exists", nil
}

func (rs RepoService) Login(credentials Credentials) (LoggedInUser, error) {
	db, err := dc.GetDB()
	if err != nil {
		return LoggedInUser{}, fmt.Errorf("failed to establish DB connection: %w", err)
	}

	var user Users
	if err := db.Model(&user).Where("username = ? AND password = ?", credentials.Username, credentials.Password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return LoggedInUser{}, common.ErrInvalidCredentials
		}
		return LoggedInUser{}, err
	}
	tokenString, err := auth.GenerateJWTToken(user.Username, user.loginId)
	if err != nil {
		fmt.Printf("Failed to generate JWT token: %v\n", err)
	}
	UserTypeInt, err := ValidateUserType(user)
	if err != nil {
		return LoggedInUser{}, common.ErrLoginInvalid
	}

	loggedUser := LoggedInUser{
		Token: tokenString,
		Details: UserDetails{
			Username: user.Username,
			UserId:   user.loginId,
			UserType: UserTypeInt,
		},
	}

	return loggedUser, nil
}

// TO CHECK THE TYPE OF EMPLOYEE
func ValidateUserType(user Users) (int, error) {
	var UserTypeInt int
	switch user.Type {
	case "employee":
		UserTypeInt = 1
	case "manager":
		UserTypeInt = 2
	case "admin":
		UserTypeInt = 3
	default:
		return UserTypeInt, errors.New("Invalid User Type")
	}
	return UserTypeInt, nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	clearCookie(w, "logintoken")
	json.NewEncoder(w).Encode("Logged Out")
}


func clearCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Set the cookie expiration time in the past
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
}

