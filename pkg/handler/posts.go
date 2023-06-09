package handler

import (
	"net/http"

	TonWork "github.com/TonWork/back"
	"github.com/gin-gonic/gin"
)

func (h *Handler) postsPOST(c *gin.Context) {
	var input TonWork.Post
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.service.Posts.Create(userId, input)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsALLGET(c *gin.Context) {
	data, err := h.service.Posts.GetAll()
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
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
