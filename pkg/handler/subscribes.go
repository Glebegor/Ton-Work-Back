package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) subscribesBuy(c *gin.Context) {
	id, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.service.Subscribes.BuySubscribe(id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
func (h *Handler) subscribesCancel(c *gin.Context) {
	id, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.service.Subscribes.CancelSubscribe(id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
func (h *Handler) subscribesTimetoend(c *gin.Context) {
	name := c.GetString("userUsername")
	id := c.GetInt("userId")
	timetoend, err := h.service.Subscribes.GetTimeToEnd(id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Username":  name,
		"TimeToEnd": timetoend,
	})
}
