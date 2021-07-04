package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/kapeel-mopkar/go-kit-demo/account/dto"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return dto.CreateUserResponse{
			Ok: ok,
		}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)
		return dto.GetUserResponse{
			Email: email,
		}, err
	}
}
