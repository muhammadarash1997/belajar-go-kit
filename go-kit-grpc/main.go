package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/muhammadarash1997/go-kit-grpc/endpoints"
	"github.com/muhammadarash1997/go-kit-grpc/pb"
	"github.com/muhammadarash1997/go-kit-grpc/services"
	"github.com/muhammadarash1997/go-kit-grpc/transports"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller")
	}

	addService := services.NewService(logger)
	addEndpoint := endpoints.MakeEndpoints(addService)
	grpcServiceServer := transports.NewGRPCServiceServer(addEndpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Create Listener
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		// Create gRPC Server
		grpcServer := grpc.NewServer()

		// Register Service Server into gRPC Server so that the gRPC Server will has services and the service will be run when gRPC be served
		pb.RegisterMathServiceServer(grpcServer, grpcServiceServer)
		level.Info(logger).Log("message", "Server started successfully")

		// Listen and Serve of Listener and gRPC Server
		grpcServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("Exit", <-errs)
}
