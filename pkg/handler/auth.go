package handler

import (
	"net/http"

	TonWork "github.com/TonWork/back"
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password_hash" binding:"required"`
}

func (h *Handler) authReg(c *gin.Context) {
	var input TonWork.User
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Authorization.CreateUser(input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func (h *Handler) authLog(c *gin.Context) {
	var input LoginUser
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password_hash)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) authProfile(c *gin.Context) {
	userParam := c.Params.ByName("Profile_Username")

	data, err := h.service.Authorization.GetUserProfile(userParam)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"username":    data.Username,
		"email":       data.Email,
		"telefon":     data.Telefon,
		"position":    data.Position,
		"description": data.Description,
		"subscribe":   data.Subscribe,
		"companies":   data.Companies,
		"name":        data.Name,
		"Surname":     data.Surname,
	})
}
