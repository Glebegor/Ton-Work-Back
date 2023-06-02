package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) postsPOST(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsALLGET(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsGET(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsPATCH(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsDELETE(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
