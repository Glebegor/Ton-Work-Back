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
	auth := router.Group("/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
		auth.POST("/profile")
	}
	apiV2 := router.Group("/api/v2/")
	{
		work := apiV2.Group("/work")
		{
			work.POST("/")
			work.GET("/")
			work.GET("/:id")
			work.PATCH("/:id")
			work.DELETE("/:id")
		}
		posts := apiV2.Group("/posts")
		{
			posts.POST("/")
			posts.GET("/")
			posts.GET("/:id")
			posts.PATCH("/:id")
			posts.DELETE("/:id")
		}
		subscribes := apiV2.Group("/subscribe")
		{
			subscribes.POST("/buy")
			subscribes.POST("/cancel")
		}
	}
	return router
}
