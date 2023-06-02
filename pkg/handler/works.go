package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) workPOST(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) workALLGET(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) workGET(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) workPATCH(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) workDELETE(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
