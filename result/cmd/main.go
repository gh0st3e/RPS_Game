package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "result/internal/proto"
	"result/internal/service"
)

const (
	Port = ":9001"
)

func main() {
	logger := logrus.New()

	resultService := service.NewService(logger)

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("[RESULT SERVICE] FAIL PORT")
	}

	resultGRPC := pb.NewResultService(resultService)

	s := grpc.NewServer()
	pb.RegisterResultServer(s, resultGRPC)

	logger.Printf("[RESULT SERVICE] gRPC Server is working on port: %s", Port)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("[RESULT SERVICE] gRPC Server was shut down")
	}
}
