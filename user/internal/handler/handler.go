package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user/internal/proto"
	"user/internal/service"
)

type Handler struct {
	service        *service.Service
	computerClient proto.ComputerClient
	resultClient   proto.ResultClient
}

func NewHandler(service *service.Service, computerClient proto.ComputerClient, resultClient proto.ResultClient) *Handler {
	return &Handler{
		service:        service,
		computerClient: computerClient,
		resultClient:   resultClient}
}

func (h *Handler) Mount(r *gin.Engine) {
	r.GET("/start/:value", h.StartGame)
}

func (h *Handler) StartGame(ctx *gin.Context) {
	userValue := ctx.Param("value")
	err := h.service.CheckValue(userValue)
	if err != nil {
		log.Printf("[USER SERVICE] Incorrect UserValue: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%s", err),
		})
		return
	}

	computerValue, err := h.computerClient.CallComputer(ctx, &proto.Empty{})
	if err != nil {
		log.Printf("[USER SERVICE] Couldn't call remote function [COMPUTER SERVICE]: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%s", err),
		})
		return
	}
	log.Printf("[USER SERVICE] ComputerValue: %s", computerValue)

	resultValue, err := h.resultClient.CallResult(ctx, &proto.RequestResult{
		UserValue:     userValue,
		ComputerValue: computerValue.Value,
	})
	if err != nil {
		log.Printf("[USER SERVICE] Couldn't call remote function [RESULT SERVICE]: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%s", err),
		})
		return
	}
	log.Printf("[USER SERVICE] Result: %s", resultValue.Winner)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User Value: %s; Computer Value: %s; Result Of The Game: %s;", userValue, computerValue.Value, resultValue.Winner),
	})
}
