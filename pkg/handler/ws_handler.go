package handler

import (
	"net/http"

	TonWork "github.com/TonWork/back"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type CreateRoomReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var input CreateRoomReq
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.hub.Rooms[input.Id] = &TonWork.Room{
		Id:      input.Id,
		Name:    input.Name,
		Clients: make(map[string]*TonWork.Client),
	}
	c.JSON(http.StatusOK, input)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	roomId := c.Param("roomId")
	clientId := c.GetString("userId")
	username := c.GetString("userUsername")

	cl := &TonWork.Client{
		Conn:     conn,
		Message:  make(chan *TonWork.Message, 10),
		Id:       clientId,
		RoomId:   roomId,
		Username: username,
	}
	m := &TonWork.Message{
		Content:  "A new user joined the room",
		RoomId:   roomId,
		Username: username,
	}
}
