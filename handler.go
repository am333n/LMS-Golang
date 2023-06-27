package main

import (
	"fmt"
	e "lms/Employees"
	login "lms/Login"
	auth "lms/Auth"
	m "lms/Managers"
	"lms/common"
	"os"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	khttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func initialiseRouter() *mux.Router {
	logger := log.NewLogfmtLogger(os.Stderr)
	controleer:=e.NewController() 
	svc := e.NewService(controleer)
	svc = e.LoggingMiddleware(logger)(svc)
	controlle:=m.NewController()
	svcc := m.NewService(controlle)
	svccc := login.RepoService{}
	//employee handlers
	opts := []khttp.ServerOption{
		khttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		khttp.ServerErrorEncoder(common.EncodeError),
		khttp.ServerBefore(auth.ExtractTokenData),
	}

	PostEmployeeHandler := httptransport.NewServer(
		e.MakePostEmployeeEndpoint(svc),
		e.DecodePostEmployeeRequest,
		e.EncodeResponse,
	)
	getEmployeesEndPoint := auth.Middleware()(e.MakeGetEmployeesEndpoint(svc))
	GetEmployeesHandler := httptransport.NewServer(
		getEmployeesEndPoint,
		e.DecodeGetEmployeesRequest,
		e.EncodeResponse,
		opts...,
	)
	MakeGetEmployeesByIdEndpoint := auth.Middleware()(e.MakeGetEmployeesByIdEndpoint(svc))
	GetEmployeeByIdHandler := httptransport.NewServer(
		MakeGetEmployeesByIdEndpoint,
		e.DecodeGetEmployeeByIdRequest,
		e.EncodeResponse,
		opts...,
	
	)
	MakeDeleteEmployeesByIdEndpoint := auth.Middleware()(e.MakeDeleteEmployeesByIdEndpoint(svc))
	DeleteEmployeeByIdHandler := httptransport.NewServer(
		MakeDeleteEmployeesByIdEndpoint,
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
		opts...,
	)
	MakeUpdateEmployeeEndpoint := auth.Middleware()(e.MakeUpdateEmployeeEndpoint(svc))
	UpdateEmployeeHandler := httptransport.NewServer(
		MakeUpdateEmployeeEndpoint,
		e.DecodeUpdateEmployee,
		e.EncodeResponse,
		opts...,
	)
	MakeApproveEmployeeEndpoint := auth.Middleware()(e.MakeApproveEmployeeEndpoint(svc))
	ApproveEmployeeHandler := httptransport.NewServer(
		MakeApproveEmployeeEndpoint,
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
		opts...,
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
	GetManageByIdHandler := httptransport.NewServer(
		m.MakeGetManagerByIdEndpoint(svcc),
		m.DecodeGetManagerByIdRequest,
		m.EncodeResponse,
	)
	DeleteManagerHandler := httptransport.NewServer(
		m.MakeDeleteManagerEndpoint(svcc),
		m.DecodeDeleteManagerRequest,
		m.EncodeResponse,
	)
	UpdateManagerHandler := httptransport.NewServer(
		m.MakeUpdateManagerEndpoint(svcc),
		m.DecodeUpdateManagerRequest,
		m.EncodeResponse,
	)
	/* ------------------------- Leave Functions handler ------------------------ */
	MakePostLeavesEndpoint:=auth.Middleware()(e.MakePostLeavesEndpoint(svc))
	PostLeavesHandler := httptransport.NewServer(
		MakePostLeavesEndpoint,
		e.DecodeGetEmployeeByIdRequest,
		e.EncodeResponse,
		opts...
	)
	DeleteLeavesHandler := httptransport.NewServer(
		e.MakeDeleteLeavesEndpoint(svc),
		e.DecodeGetEmployeeByIdRequest,
		e.EncodeResponse,
	)
	EnterLeaveHandler := httptransport.NewServer(
		e.MakeEnterLeaveEndpoint(svc),
		e.DecodeEnterLeaveRequest,
		e.EncodeResponse,
	)
	/* ---------------------------- Request Handlers ---------------------------- */
	MakePostLeaveRequestEndpoint:=auth.Middleware()(e.MakePostLeaveRequestEndpoint(svc))
	PostLeaveRequestHandler := httptransport.NewServer(
		MakePostLeaveRequestEndpoint,
		e.DecodePostLeaveRequest,
		e.EncodeResponse,
		opts...,
	)
	GetRequestHandler := httptransport.NewServer(
		e.MakeGetRequestEndpoint(svc),
		e.DecodeGetRequestRequest,
		e.EncodeResponse,
	)
	ApproveLeaveRequestHandler := httptransport.NewServer(
		e.MakeApproveLeaveRequestEndpoint(svc),
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
	)
	DeleteRequestHandler := httptransport.NewServer(
		e.MakeDeleteLeaveRequestEndpoint(svc),
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
	)
	GetLeavesHandler := httptransport.NewServer(
		e.MakeGetLeavesEndpoint(svc),
		e.DecodeGetEmployeesRequest,
		e.EncodeResponse,
	)
	GetLeavesByIdHandler := httptransport.NewServer(
		e.MakeGetLeavesByIdEndpoint(svc),
		e.DecodeDeleteEmployeeByIdRequest,
		e.EncodeResponse,
	)
	/* ----------------------------- Login/Signup handler ----------------------------- */

	Signuphandler := httptransport.NewServer(
		login.MakeSignupEndpoint(svccc),
		login.DecodeSignupRequest,
		e.EncodeResponse,
	)
	LoginHandler := httptransport.NewServer(
		login.MakeLoginEndpoint(svccc),
		login.DecodeLoginRequest,
		login.EncodeLoginResponse,
	)

	/* ------------------------------ router setup ------------------------------ */
	router := mux.NewRouter()
	//Login
	router.Handle("/login", LoginHandler).Methods("POST")
	router.Handle("/signup/{type}", Signuphandler).Methods("POST")
	router.HandleFunc("/logout", login.Logout).Methods("POST")

	//employee
	router.Handle("/Employees", GetEmployeesHandler).Methods("GET")
	router.Handle("/Employees", PostEmployeeHandler).Methods("POST")
	router.Handle("/Employees/{id}", GetEmployeeByIdHandler).Methods("GET")
	router.Handle("/Employees/{id}", DeleteEmployeeByIdHandler).Methods("DELETE")
	router.Handle("/Employees/{id}", UpdateEmployeeHandler).Methods("PUT")
	router.Handle("/Employees/{id}", ApproveEmployeeHandler).Methods("PATCH")

	//EmployeeLeave
	router.Handle("/Employees/Leave/", GetLeavesHandler).Methods("GET")
	router.Handle("/Employees/leave/{id}", PostLeavesHandler).Methods("POST")
	router.Handle("/Employees/leave/{id}", DeleteLeavesHandler).Methods("DELETE")
	router.Handle("/Employees/leave/{id}", EnterLeaveHandler).Methods("PUT")
	router.Handle("/Employees/leave/{id}", GetLeavesByIdHandler).Methods("GET")

	//Request
	router.Handle("/Employees/leave/request/", PostLeaveRequestHandler).Methods("POST")
	router.Handle("/Employees/leave/request/", GetRequestHandler).Methods("GET")
	router.Handle("/Employees/leave/request/{id}", ApproveLeaveRequestHandler).Methods("PATCH")
	router.Handle("/Employees/leave/request/{id}", DeleteRequestHandler).Methods("DELETE")

	//manager
	router.Handle("/Managers", GetManagersHandler).Methods("GET")
	router.Handle("/Managers", PostManagerHandler).Methods("POST")
	router.Handle("/Managers/{id}", GetManageByIdHandler).Methods("GET")
	router.Handle("/Managers/{id}", DeleteManagerHandler).Methods("DELETE")
	router.Handle("/Managers/{id}", UpdateManagerHandler).Methods("PUT")

	fmt.Println("The server is running on 8080")

	return router
}
