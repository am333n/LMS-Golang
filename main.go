package main

import (
	"fmt"
	dc "lms/Database"
	e "lms/Employees"
	l "lms/Login"
	m "lms/Managers"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func initialiseRouter() *mux.Router {
	svc := e.RepoService{}
	svcc := m.RepoService{}
	PostEmployeeHandler := httptransport.NewServer(
		e.MakePostEmployeeEndpoint(svc),
		e.DecodePostEmployeeRequest,
		e.EncodeResponse,
	)
	GetEmployeesHandler := httptransport.NewServer(
		e.MakeGetEmployeesEndpoint(svc),
		e.DecodeGetEmployeesRequest,
		e.EncodeResponse,
	)
	GetEmployeeByIdHandler := httptransport.NewServer(
		e.MakeGetEmployeesByIdEndpoint(svc),
		e.DecodeGetEmployeeByIdRequest,
		e.EncodeResponse,
	)
	DeleteEmployeeByIdHandler := httptransport.NewServer(
		e.MakeDeleteEmployeesByIdEndpoint(svc),
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
	)
	UpdateEmployeeHandler := httptransport.NewServer(
		e.MakeUpdateEmployeeEndpoint(svc),
		e.DecodeUpdateEmployee,
		e.EncodeResponse,
	)
	GetManagersHandler := httptransport.NewServer(
		m.MakeGetManagersEndpoint(svcc),
		m.DecodeGetManagersRequest,
		m.EncodeResponse,
	)
	PostManagerHandler := httptransport.NewServer(
		m.MakePostManagerEndpoint(svcc),
		m.DecodePostManagerRequest,
		m.EncodeResponse,
	)
	GetManageByIdHandler:=httptransport.NewServer(
		m.MakeGetManagerByIdEndpoint(svcc),
        m.DecodeGetManagerByIdRequest,
        m.EncodeResponse,
	)
	router := mux.NewRouter()
	//Login
	router.HandleFunc("/login", l.Login).Methods("POST")
	router.HandleFunc("/signup", l.Signup).Methods("POST")
	//employee
	router.Handle("/Employees", GetEmployeesHandler).Methods("GET")
	router.Handle("/Employees", PostEmployeeHandler).Methods("POST")
	router.Handle("/Employees/{id}", GetEmployeeByIdHandler).Methods("GET")
	router.Handle("/Employees/{id}", DeleteEmployeeByIdHandler).Methods("DELETE")
	router.Handle("/Employees/{id}", UpdateEmployeeHandler).Methods("PUT")
	//manager
	router.Handle("/Managers", GetManagersHandler).Methods("GET")
	router.Handle("/Managers", PostManagerHandler).Methods("POST")
	router.Handle("/Managers/{id}",GetManageByIdHandler).Methods("GET")
	// router.HandleFunc("/Managers/{id}", m.DeleteManager).Methods("DELETE")
	// router.HandleFunc("/Managers/{id}", m.PutManager).Methods("PUT")
	fmt.Println("The server is running on 8080")

	return router
}

func main() {
	_, err := dc.GetDB()
	if err != nil {
		panic(err)
	}

	dc.DB.AutoMigrate(&e.Employees{}, &m.Manager{}, &l.Users{})
	router := initialiseRouter()
	defer http.ListenAndServe(":8080", router)

}
