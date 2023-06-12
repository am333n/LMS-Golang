package Employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

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
func  MakeGetEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:=request.(int)
		v,err:=svc.GetEmployeeById(req)
		if err!=nil{
            return GetEmployeeByIdResponse{v,err.Error()},nil
        }
		return GetEmployeeByIdResponse{v,""},nil

	}
}
func MakeDeleteEmployeesByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:=request.(int)
		result,err:=svc.DeleteEmployeeById(req)
		if err!=nil{
			return DeleteEmployeeByIdResponse{result,err.Error()},nil
		}
		return DeleteEmployeeByIdResponse{result,""},nil
	}

}
func MakeUpdateEmployeeEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:=request.(UpdateEmployeeRequest)
		result,res,err:=svc.UpdateEmployee(req.id,req.employee)
		if err!=nil{
            return UpdateEmployeeResponse{result,res,err.Error()},nil
        }
		return UpdateEmployeeResponse{result,res,""},nil

	}
}
