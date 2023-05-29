package handler

import (
	service "github.com/TonWork/back/pkg/service"
	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	apiV2 := router.Group("/api/v2/")
	{
		auth := apiV2.Group("/auth")
		{

		}
	}
	return router
}
