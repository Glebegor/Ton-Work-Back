package handler

import (
	"errors"
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

func GetUserById(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		newResponse(c, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("User id is not found")
	}
	idint, ok := id.(int)
	if !ok {
		newResponse(c, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("User id is not found")
	}
	return idint, nil
}
