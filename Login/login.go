package login

import (
	"encoding/json"
	"fmt"
	dc "lms/Database"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var jwtKey = []byte("secretkey")

type Users struct {
	gorm.Model
	loginId  int    `gorm:"autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func SignupEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var user Users
	json.NewDecoder(r.Body).Decode(&user)
	user.Type = params["type"]
	if user.Type == "employee" {
		
	}
	if user.Type == "admin" {

	}
	dc.DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func getLoginDetails(w http.ResponseWriter, r *http.Request) (Details, Credentials,error) {
	w.Header().Set("Content-Type", "application/json")
	var credentials Credentials
	json.NewDecoder(r.Body).Decode(&credentials)
	var user Users
	var details Details
	db,err:=dc.GetDB()
	if err!=nil{
		return details,credentials, fmt.Errorf("%w",err)
	}
	if err := db.Where("username=?", credentials.Username).First(&user).Error; err != nil {
		fmt.Printf("%s", "cannot get details")
	}
	details.Id = user.loginId
	details.Username = user.Username
	details.Password = user.Password
	// json.NewEncoder(w).Encode(&details)
	// json.NewEncoder(w).Encode(&credentials)
	return details, credentials,nil

}

type Details struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var details Details

func Login(w http.ResponseWriter, r *http.Request) {
	details, credentials,err:= getLoginDetails(w, r)
	expectedPassword := credentials.Password
	if expectedPassword != details.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	experationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Id:       details.Id,
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: experationTime,
	})
}
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte("hello world"))

}
