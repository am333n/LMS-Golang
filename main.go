package main

import (
	dc "lms/Database"
	e "lms/Employees"
	login "lms/Login"
	m "lms/Managers"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	_, err := dc.GetDB()
	if err != nil {
		panic(err)
	}
	dc.DB.AutoMigrate(&e.Employees{}, &m.Manager{}, &login.Users{}, e.Leaves{}, e.Requests{})
	router := initialiseRouter()
	defer http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET","POST","DELETE","PATCH","PUT","OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"}),
		handlers.AllowCredentials(),

	)(router))


}
