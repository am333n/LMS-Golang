package managers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetManagersEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetManagers()
		if err != nil {
			return GetManagersResponse{v, err.Error()}, nil
		}
		return GetManagersResponse{v, ""}, nil
	}
}
func MakePostManagerEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PostManagerRequest)
		v, err := svc.PostManager(req.manager)
		if err != nil {
			return PostManagerResponse{v, err.Error()}, nil
		}
		return PostManagerResponse{v, ""}, nil

	}
}
func MakeGetManagerByIdEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:= request.(int)
		v,err:=svc.GetManagerById(req)
		if err!=nil{
			return GetManagerByIdResponse{v,err.Error()},nil
		}
		return GetManagerByIdResponse{v,""},nil
    }	
}
func MakeDeleteManagerEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:=request.(int)
		v,err:=svc.DeleteManager(req)
		if err!=nil{
			return DeleteManagerResponse{v,err.Error()},nil
		}
		return DeleteManagerResponse{v,""},nil
	}
}
func MakeUpdateManagerEndpoint(svc Service) endpoint.Endpoint{
	return func(_ context.Context, request interface{}) (interface{}, error){
		req:=request.(UpdateManagerRequest)
		v,err:=svc.UpdateManager(req.id,req.manager)
		if err!=nil{
			return UpdateManagerResponse{v,err.Error()},nil
		}
		return UpdateManagerResponse{v,""},nil
	}
	
}

