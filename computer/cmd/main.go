package main

import (
	pb "computer/internal/proto"
	"computer/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	Port = ":9000"
)

func main() {
	logger := logrus.New()
	computerService := service.NewService(logger)

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		logger.Fatalf("[COMPUTER SERVICE] FAIL PORT")
	}

	computerGRPC := pb.NewComputerService(computerService)

	s := grpc.NewServer()
	pb.RegisterComputerServer(s, computerGRPC)

	logger.Printf("[COMPUTER SERVICE] gRPC Server is working on port: %s", Port)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("[COMPUTER SERVICE] gRPC Server was shut down")
	}

}
