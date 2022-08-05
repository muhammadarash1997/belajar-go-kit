package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/go-kit-http/models"
	"github.com/muhammadarash1997/go-kit-http/services"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.CreateUserRequest)

		message, err := s.CreateUser(ctx, req.Email, req.Password)
		return models.CreateUserResponse{
			Message: message,
		}, err
	}
}

func makeGetUserEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.GetUserRequest)

		email, err := s.GetUser(ctx, req.Id)
		return models.GetUserResponse{
			Email: email,
		}, err
	}
}
