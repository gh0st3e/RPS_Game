package proto

import (
	"context"
	"result/internal/service"
)

type ResultService struct {
	service *service.Service
	UnimplementedResultServer
}

func NewResultService(service *service.Service) *ResultService {
	return &ResultService{service: service}
}

func (r *ResultService) CallResult(ctx context.Context, in *RequestResult) (*ResponseResult, error) {
	winner := r.service.WhoIsWinner(in.UserValue, in.ComputerValue)
	return &ResponseResult{Winner: winner}, nil
}
