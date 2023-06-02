package handler

import (
	"net/http"

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
		auth.POST("/register", h.HandlerEmptyFunc)
		auth.POST("/login", h.HandlerEmptyFunc)
		auth.POST("/profile", h.HandlerEmptyFunc)
	}
	apiV2 := router.Group("/api/v2/")
	{
		work := apiV2.Group("/work")
		{
			work.POST("/", h.HandlerEmptyFunc)
			work.GET("/", h.HandlerEmptyFunc)
			work.GET("/:id", h.HandlerEmptyFunc)
			work.PATCH("/:id", h.HandlerEmptyFunc)
			work.DELETE("/:id", h.HandlerEmptyFunc)
		}
		posts := apiV2.Group("/posts")
		{
			posts.POST("/", h.HandlerEmptyFunc)
			posts.GET("/", h.HandlerEmptyFunc)
			posts.GET("/:id", h.HandlerEmptyFunc)
			posts.PATCH("/:id", h.HandlerEmptyFunc)
			posts.DELETE("/:id", h.HandlerEmptyFunc)
		}
		subscribes := apiV2.Group("/subscribe")
		{
			subscribes.POST("/buy", h.HandlerEmptyFunc)
			subscribes.POST("/cancel", h.HandlerEmptyFunc)
		}
	}
	return router
}

func (h *Handler) HandlerEmptyFunc(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
