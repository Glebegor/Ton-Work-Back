package handler

import (
	service "github.com/TonWork/back/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

// func (h *Handler) InitRoutes() *gin.Engine {
// 	return
// }
