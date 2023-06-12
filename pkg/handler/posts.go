package handler

import (
	"net/http"
	"strconv"

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
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newResponse(c, http.StatusNotFound, err.Error())
		return
	}
	data, err := h.service.Posts.GetById(idInt)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})

}
func (h *Handler) postsPUT(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	var input TonWork.PostUpdate
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Posts.Update(idInt, input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) postsDELETE(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.service.Posts.Delete(idInt); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
