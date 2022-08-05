package transports

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/muhammadarash1997/go-kit-grpc/endpoints"
	"github.com/muhammadarash1997/go-kit-grpc/models"
	"github.com/muhammadarash1997/go-kit-grpc/pb"
)

type GRPCServiceServer struct {
	pb.UnimplementedMathServiceServer
	add grpc.Handler
}

func NewGRPCServiceServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &GRPCServiceServer{
		add: grpc.NewServer(
			endpoints.Add,      // Akan dijalankan setelah decodeMathRequest dijalankan
			decodeMathRequest,  // Akan dijalankan pertama kali dan bertugas mengekstrak payload bertipe pb.Request menjadi payload bertipe interface{}
			encodeMathResponse, // Akan dijalankan terakhir yang mana bertugas mengekstrak data bertipe interface{} menjadi data bertipe pb.Response dan bertugas untuk mengembalikan response ke client
		),
	}
}

func (s *GRPCServiceServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	// Setelah s.add.ServeGRPC maka urutan fungsi yang akan dijalankan seperti ini :
	// - decodeMathRequest
	// - endpoints.Add
	// - encodeMathResponse
	_, res, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.MathResponse), nil
}

// decodeMathRequest untuk mengubah payload bertipe pb (yang diterima dari client) ke struct model yang selanjutnya akan dipakai dan diproses oleh endpoints
func decodeMathRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return models.MathRequest{
		NumA: req.NumA,
		NumB: req.NumB,
	}, nil
}

// encodeMathResponse untuk mengubah data yang telah diproses oleh endpoints (yang bertipe struct) ke pb dan akan dikembalikan ke s.add.ServeGRPC
func encodeMathResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(models.MathResponse)
	return &pb.MathResponse{
		Result: res.Result,
	}, nil
}
