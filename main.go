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
	//employee handlers
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
	ApproveEmployeeHandler := httptransport.NewServer(
		e.MakeApproveEmployeeEndpoint(svc),
        e.DecodeDeleteEmployeeByIdRequest,
        e.EncodeResponse,
	)
	//manager handlers
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
	DeleteManagerHandler:=httptransport.NewServer(
		m.MakeDeleteManagerEndpoint(svcc),
		m.DecodeDeleteManagerRequest,
		m.EncodeResponse,
	)
	UpdateManagerHandler:=httptransport.NewServer(
		m.MakeUpdateManagerEndpoint(svcc),
		m.DecodeUpdateManagerRequest,
		m.EncodeResponse,
	)
	/* ------------------------- Leave Functions handler ------------------------ */
	PostLeavesHandler:=httptransport.NewServer(
		e.MakePostLeavesEndpoint(svc),
        e.DecodeGetEmployeeByIdRequest,
        e.EncodeResponse,
	)
	DeleteLeavesHandler:=httptransport.NewServer(
		e.MakeDeleteLeavesEndpoint(svc),
        e.DecodeGetEmployeeByIdRequest,
        e.EncodeResponse,
	)
	EnterLeaveHandler:=httptransport.NewServer(
		e.MakeEnterLeaveEndpoint(svc),
        e.DecodeEnterLeaveRequest,
        e.EncodeResponse,
	)
	/* ---------------------------- Request Handlers ---------------------------- */
	PostLeaveRequestHandler:=httptransport.NewServer(
		e.MakePostLeaveRequestEndpoint(svc),
        e.DecodePostLeaveRequest,
        e.EncodeResponse,
	)

	/* ------------------------------ router setup ------------------------------ */
	router := mux.NewRouter()
	//Login
	router.HandleFunc("/login", l.Login).Methods("POST")
	router.HandleFunc("/signup", l.SignupEmployee).Methods("POST")


	//employee
	router.Handle("/Employees", GetEmployeesHandler).Methods("GET")
	router.Handle("/Employees", PostEmployeeHandler).Methods("POST")
	router.Handle("/Employees/{id}", GetEmployeeByIdHandler).Methods("GET")
	router.Handle("/Employees/{id}", DeleteEmployeeByIdHandler).Methods("DELETE")
	router.Handle("/Employees/{id}", UpdateEmployeeHandler).Methods("PUT")
	router.Handle("/Employees/{id}", ApproveEmployeeHandler).Methods("PATCH")
	

	//EmployeeLeave
	router.Handle("/Employees/leave/{id}", PostLeavesHandler).Methods("POST")
	router.Handle("/Employees/leave/{id}", DeleteLeavesHandler).Methods("DELETE")
	router.Handle("/Employees/leave/{id}",EnterLeaveHandler).Methods("PUT")


	//Request
	router.Handle("/Employees/leave/request/",PostLeaveRequestHandler ).Methods("POST")


	//manager
	router.Handle("/Managers", GetManagersHandler).Methods("GET")
	router.Handle("/Managers", PostManagerHandler).Methods("POST")
	router.Handle("/Managers/{id}",GetManageByIdHandler).Methods("GET")
	router.Handle("/Managers/{id}", DeleteManagerHandler).Methods("DELETE")
	router.Handle("/Managers/{id}",UpdateManagerHandler).Methods("PUT")
	fmt.Println("The server is running on 8080")

	return router
}

func main() {
	_, err := dc.GetDB()
	if err != nil {
		panic(err)
	}

	dc.DB.AutoMigrate(&e.Employees{}, &m.Manager{}, &l.Users{},e.Leaves{},e.Requests{})
	router := initialiseRouter()
	defer http.ListenAndServe(":8080", router)

}
