package Employee

import (
	"context"
	auth "lms/Auth"

	"github.com/go-kit/kit/endpoint"
)

/* --------------------------- Endpoint Connection -------------------------- */

type Endpoints struct {
	PostEmployeeEndpoint        endpoint.Endpoint
	GetEmployeesEndpoint        endpoint.Endpoint
	DeleteEmployeesByIdEndpoint endpoint.Endpoint
	UpdateEmployeeEndpoint      endpoint.Endpoint
	ApproveEmployeeEndpoint     endpoint.Endpoint
	GetEmployeesByIdEndpoint    endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostEmployeeEndpoint:        auth.Middleware()(MakePostEmployeeEndpoint(s)),
		GetEmployeesEndpoint:        auth.Middleware()(MakeGetEmployeesEndpoint(s)),
		DeleteEmployeesByIdEndpoint: auth.Middleware()(MakeDeleteEmployeesByIdEndpoint(s)),
		UpdateEmployeeEndpoint:      auth.Middleware()(MakeUpdateEmployeeEndpoint(s)),
		ApproveEmployeeEndpoint:     auth.Middleware()(MakeApproveEmployeeEndpoint(s)),
		GetEmployeesByIdEndpoint:    auth.Middleware()(MakeGetEmployeesByIdEndpoint(s)),
	}
}

/* ------------------------ Employee CRUD operations ------------------------ */

func MakePostEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostEmployeeRequest)
		v, err := svc.PostEmployee(req.employee)
		if err != nil {
			return PostEmployeeResponse{v, err.Error()}, nil
		}
		return PostEmployeeResponse{v, ""}, nil
	}
}
func MakeGetEmployeesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(GetEmployeesRequest)
		v, err := svc.GetEmployees(ctx)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
func MakeGetEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.GetEmployeeById(req)
		if err != nil {
			return GetEmployeeByIdResponse{v, err.Error()}, nil
		}
		return GetEmployeeByIdResponse{v, ""}, nil
	}
}
func MakeDeleteEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		result, err := svc.DeleteEmployeeById(ctx,req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}
}
func MakeUpdateEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmployeeRequest)
		result, res, err := svc.UpdateEmployee(ctx,req.id, req.employee)
		if err != nil {
			return UpdateEmployeeResponse{result, res, err.Error()}, nil
		}
		return UpdateEmployeeResponse{result, res, ""}, nil
	}
}
func MakeApproveEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		res, err := svc.ApproveEmployee(ctx,req)
		if err != nil {
			return DeleteEmployeeByIdResponse{res, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{res, ""}, nil
	}
}

/* ------------------------- Employee Leave Function ------------------------ */

func MakePostLeavesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		result, err := svc.PostLeaves(ctx,req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}
}
func MakeDeleteLeavesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		result, err := svc.DeleteLeaves(ctx,req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}
}

func MakeEnterLeaveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnterLeaveRequest)
		v, err := svc.EnterLeaves(ctx,req.id, req.leave)
		if err != nil {
			return EnterLeaveResponse{v, err.Error()}, nil
		}
		return EnterLeaveResponse{v, ""}, nil
	}
}
func MakeGetLeavesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetLeaves(ctx)
		if err != nil {
			return EnterLeaveResponse{v, err.Error()}, nil
		}
		return EnterLeaveResponse{v, ""}, nil
	}
}

/* ---------------------------- Request Endpoints --------------------------- */

func MakePostLeaveRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostLeaveRequestRequest)
		v, err := svc.PostLeaveRequest(ctx,req.request)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeGetRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetLeaveRequest(ctx)
		if err != nil {
			return GetLeaveRequestResponse{v, err.Error()}, nil
		}
		return GetLeaveRequestResponse{v, ""}, nil
	}
}
func MakeApproveLeaveRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.ApproveLeaveRequest(ctx,req)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeDeleteLeaveRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.DeleteLeaveRequest(ctx,req)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeGetLeavesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.GetLeavesById(ctx,req)
		if err != nil {
			return GetLeavesByIdResponse{v, err.Error()}, nil
		}
		return GetLeavesByIdResponse{v, ""}, nil
	}
}
