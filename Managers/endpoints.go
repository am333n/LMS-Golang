package managers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetManagersEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v,err:=svc.GetManagers()
		if err!=nil{
				return GetManagersResponse{v,err.Error()},nil
		}
		return GetManagersResponse{v,""},nil
	}
}
func MakePostManagerEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req:=request.(PostManagerRequest)
		v,err:=svc.PostManager(req.status,req.manager)
		if err!=nil{
			return PostManagerResponse{"",v,err.Error()},nil
		}
		return PostManagerResponse{"Manager Added",v,""},nil

	}
}
