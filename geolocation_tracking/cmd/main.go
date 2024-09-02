package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
)

type message struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Hub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan message
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn] = true
		case conn := <-h.Unregister:
			delete(h.Clients, conn)
		case msg := <-h.Broadcast:
			for conn := range h.Clients {
				_ = conn.WriteJSON(msg)
			}
		}
	}
}

func main() {
	h := &Hub{
		Broadcast:  make(chan message),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Clients:    make(map[*websocket.Conn]bool),
	}

	go h.Run()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Use("/ws/location", websocket.New(LocationTracking(h)))

	err := app.Listen("127.0.0.1:8000")

	if err != nil {
		log.Fatal(err)
	}
}

func LocationTracking(h *Hub) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		defer func() {
			h.Unregister <- c
			//_ = c.Close()
		}()

		//name := c.Query("name", "driver")
		h.Register <- c

		for {
			var msg message
			err := c.ReadJSON(&msg)

			if err != nil {
				log.Fatal(err)
				return
			}

			h.Broadcast <- msg
		}
	}
}
