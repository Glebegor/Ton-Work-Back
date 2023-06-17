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
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

type RoomRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			Id:   r.Id,
			Name: r.Name,
		})
	}
	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients := make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			Id:       c.Id,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
