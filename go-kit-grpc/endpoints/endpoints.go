package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadarash1997/go-kit-grpc/models"
	"github.com/muhammadarash1997/go-kit-grpc/services"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoints(s),
	}
}

func makeAddEndpoints(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.MathRequest)

		result, err := s.Add(ctx, req.NumA, req.NumB)

		return models.MathResponse{
			Result: result,
		}, err
	}
}
