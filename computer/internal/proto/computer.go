package proto

import (
	"computer/internal/service"
	"context"
)

type ComputerService struct {
	service *service.Service
	UnimplementedComputerServer
}

func NewComputerService(service *service.Service) *ComputerService {
	return &ComputerService{service: service}

}

func (c *ComputerService) CallComputer(ctx context.Context, in *Empty) (*ResponseComputer, error) {
	value := c.service.GenerateValue()
	return &ResponseComputer{
		Value: value,
	}, nil
}
