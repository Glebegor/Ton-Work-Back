package TonWork

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	Id       string `json:"id"`
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}
