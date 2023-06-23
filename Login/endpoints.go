package login

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeSignupEndpoint(svc LoginService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignupRequest)
		v, err := svc.Signup(req.TypeId, req.user)
		if err != nil {
			return SignupResponse{"", err.Error()}, nil
		}
		return SignupResponse{v, ""}, nil
	}
}


func MakeLoginEndpoint(svc LoginService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req:=request.(LoginRequest)
		v,err:=svc.Login(req.credentials)
		if err!=nil{
			return LoginResponse{v,err.Error()},nil
		}
		return LoginResponse{v,""},nil
	}
}