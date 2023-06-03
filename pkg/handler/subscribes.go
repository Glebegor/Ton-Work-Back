package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) subscribesBuy(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
func (h *Handler) subscribesCancel(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
