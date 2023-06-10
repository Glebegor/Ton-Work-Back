package handler

import (
	"net/http"
	"strconv"

	TonWork "github.com/TonWork/back"
	"github.com/gin-gonic/gin"
)

func (h *Handler) workPOST(c *gin.Context) {
	var input TonWork.Work
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.service.Work.Create(userId, input)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
func (h *Handler) workALLGET(c *gin.Context) {
	data, err := h.service.Work.GetAll()
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (h *Handler) workGET(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newResponse(c, http.StatusNotFound, err.Error())
		return
	}
	data, err := h.service.Work.GetById(idInt)
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
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
