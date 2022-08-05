package transports

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/muhammadarash1997/go-kit-http/endpoints"
	"github.com/muhammadarash1997/go-kit-http/models"
	"github.com/muhammadarash1997/go-kit-http/pb"
)

type GRPCServiceServer struct {
	pb.UnimplementedUserServiceServer
	createUser grpc.Handler
	getUser    grpc.Handler
}

func NewGRPCServiceServer(endpoints endpoints.Endpoints, logger log.Logger) pb.UserServiceServer {
	return &GRPCServiceServer{
		createUser: grpc.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
		getUser: grpc.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),
	}
}

func (s *GRPCServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, res, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateUserResponse), nil
}

func decodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return models.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func encodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(models.CreateUserResponse)
	return &pb.CreateUserResponse{
		Message: res.Message,
	}, nil
}

func (s *GRPCServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, res, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUserResponse), nil
}

func decodeGetUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return models.GetUserRequest{
		Id: req.Id,
	}, nil
}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(models.GetUserResponse)
	return &pb.GetUserResponse{
		Email: res.Email,
	}, nil
}
