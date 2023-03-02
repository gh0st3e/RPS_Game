package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user/internal/handler"
	"user/internal/proto"
	"user/internal/service"
)

const (
	computerPort = "computer:9000"
	resultPort   = "result:9001"
)

func main() {
	//computerConn, err := grpc.Dial(computerPort, grpc.WithInsecure(), grpc.WithNoProxy())
	logger := logrus.New()

	userService := service.NewService(logger)

	computerConn, err := grpc.Dial(computerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't dial [COMPUTER SERVICE] throug GRPC")
	}
	defer computerConn.Close()
	computerClient := proto.NewComputerClient(computerConn)

	resultConn, err := grpc.Dial(resultPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't dial [RESULT SERVICE] throug GRPC")
	}
	defer resultConn.Close()
	resultClient := proto.NewResultClient(resultConn)

	userHandler := handler.NewHandler(userService, computerClient, resultClient)
	server := gin.New()
	userHandler.Mount(server)
	server.Run(":8081")
}
