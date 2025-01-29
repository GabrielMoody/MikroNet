package dto

import (
	"github.com/gofiber/contrib/websocket"
	"sync"
)

type Message struct {
	UserID string  `json:"user_id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

type Room struct {
	ID      string
	Clients map[*websocket.Conn]bool
	Mutex   sync.Mutex
}

type Hub struct {
	Rooms      map[string]*Room
	Msg        chan Message
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}
