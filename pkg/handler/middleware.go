package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) Indentification(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == " " {
		newResponse(c, http.StatusUnauthorized, "Empty auth header.")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newResponse(c, http.StatusUnauthorized, "Invalid Auth header.")
		return
	}
	userId, userUsername, userName, userSurname, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
	c.Set("userUsername", userUsername)
	c.Set("userName", userName)
	c.Set("userSurname", userSurname)
}
