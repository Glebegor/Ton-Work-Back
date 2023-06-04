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
	router.Use(CORS())
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.authReg)
		auth.POST("/login", h.authLog)
		auth.GET("/profile/:Profile_Username", h.authProfile)
	}
	apiV2 := router.Group("/api/v2/")
	{
		work := apiV2.Group("/work")
		{
			work.POST("/", h.workPOST)
			work.GET("/", h.workALLGET)
			work.GET("/:id", h.workGET)
			work.PATCH("/:id", h.workPATCH)
			work.DELETE("/:id", h.workDELETE)
		}
		posts := apiV2.Group("/posts")
		{
			posts.POST("/", h.postsPOST)
			posts.GET("/", h.postsALLGET)
			posts.GET("/:id", h.postsGET)
			posts.PATCH("/:id", h.postsPATCH)
			posts.DELETE("/:id", h.postsDELETE)
		}
		subscribes := apiV2.Group("/subscribe")
		{
			subscribes.POST("/buy", h.subscribesBuy)
			subscribes.POST("/cancel", h.subscribesCancel)
		}
	}
	return router
}
