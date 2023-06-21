package main

import (
	dc "lms/Database"
	e "lms/Employees"
	login "lms/Login"
	m "lms/Managers"
	"net/http"
)

func main() {
	_, err := dc.GetDB()
	if err != nil {
		panic(err)
	}
	dc.DB.AutoMigrate(&e.Employees{}, &m.Manager{}, &login.Users{}, e.Leaves{}, e.Requests{})
	router := initialiseRouter()
	defer http.ListenAndServe(":8080", router)

}
