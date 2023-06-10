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
			noIndentification := work.Group("/")
			{
				noIndentification.GET("", h.workALLGET)
				noIndentification.GET(":id", h.workGET)
			}
			Indentification := work.Group("/", h.Indentification)
			{
				Indentification.POST("", h.workPOST)
				Indentification.PUT(":id", h.workPUT)
				Indentification.DELETE(":id", h.workDELETE)
			}

		}

		posts := apiV2.Group("/posts")
		{
			noIndentification := posts.Group("/")
			{
				noIndentification.GET("", h.postsALLGET)
				noIndentification.GET(":id", h.postsGET)
			}
			Indentification := posts.Group("/", h.Indentification)
			{
				Indentification.POST("", h.postsPOST)
				Indentification.PUT(":id", h.postsPUT)
				Indentification.DELETE(":id", h.postsDELETE)
			}
		}
		subscribes := apiV2.Group("/subscribe")
		{
			Indentification := subscribes.Group("/", h.Indentification)
			{
				Indentification.POST("buy", h.subscribesBuy)
				Indentification.POST("cancel", h.subscribesCancel)
			}
		}
	}
	return router
}
