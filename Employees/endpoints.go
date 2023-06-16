package Employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/* ------------------------ Employee CRUD operations ------------------------ */

func MakePostEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PostEmployeeRequest)
		v, err := svc.PostEmployee(req.employee)
		if err != nil {
			return PostEmployeeResponse{v, err.Error()}, nil
		}
		return PostEmployeeResponse{v, ""}, nil
	}
}
func MakeGetEmployeesEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//req := request.(GetEmployeesRequest)
		v, err := svc.GetEmployees()
		if err != nil {
			return GetEmployeesResponse{v, err.Error()}, nil
		}
		return GetEmployeesResponse{v, ""}, nil
	}
}
func MakeGetEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.GetEmployeeById(req)
		if err != nil {
			return GetEmployeeByIdResponse{v, err.Error()}, nil
		}
		return GetEmployeeByIdResponse{v, ""}, nil

	}
}
func MakeDeleteEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		result, err := svc.DeleteEmployeeById(req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}

}
func MakeUpdateEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmployeeRequest)
		result, res, err := svc.UpdateEmployee(req.id, req.employee)
		if err != nil {
			return UpdateEmployeeResponse{result, res, err.Error()}, nil
		}
		return UpdateEmployeeResponse{result, res, ""}, nil

	}
}
func MakeApproveEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		res, err := svc.ApproveEmployee(req)
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
		result, err := svc.PostLeaves(req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}
}
func MakeDeleteLeavesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		result, err := svc.DeleteLeaves(req)
		if err != nil {
			return DeleteEmployeeByIdResponse{result, err.Error()}, nil
		}
		return DeleteEmployeeByIdResponse{result, ""}, nil
	}
}

func MakeEnterLeaveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnterLeaveRequest)
		v, err := svc.EnterLeaves(req.id, req.leave)
		if err != nil {
			return EnterLeaveResponse{v, err.Error()}, nil
		}
		return EnterLeaveResponse{v, ""}, nil
	}
}
func MakeGetLeavesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetLeaves()
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
		v, err := svc.PostLeaveRequest(req.request)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeGetRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetLeaveRequest()
		if err != nil {
			return GetLeaveRequestResponse{v, err.Error()}, nil
		}
		return GetLeaveRequestResponse{v, ""}, nil
	}
}
func MakeApproveLeaveRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.ApproveLeaveRequest(req)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeDeleteLeaveRequestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.DeleteLeaveRequest(req)
		if err != nil {
			return PostLeaveRequestResponse{v, err.Error()}, nil
		}
		return PostLeaveRequestResponse{v, ""}, nil
	}
}
func MakeGetLeavesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		v, err := svc.GetLeavesById(req)
		if err != nil {
			return GetLeavesByIdResponse{v, err.Error()}, nil
		}
		return GetLeavesByIdResponse{v, ""}, nil
	}
}
