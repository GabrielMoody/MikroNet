package ws

import (
	"context"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/helper"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/midleware"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/gofiber/contrib/websocket"
	"log"
	"os"
	"sync"
)

type Room struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var rooms = make(map[string]*Room)
var roomsMu sync.Mutex

type GeoTrack interface {
	LocationTracking(ctx context.Context) func(*websocket.Conn)
	Run()
}

type GeoTrackImpl struct {
	h    *dto.Hub
	repo repository.GeoTrackRepository
}

func (a *GeoTrackImpl) Run() {
	for {
		select {
		case conn := <-a.h.Register:
			a.h.Clients[conn] = true
		case conn := <-a.h.Unregister:
			delete(a.h.Clients, conn)
		case msg := <-a.h.Broadcast:
			for conn := range a.h.Clients {
				_ = conn.WriteJSON(msg)
			}
		}
	}
}

func (a *GeoTrackImpl) LocationTracking(ctx context.Context) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		defer func() {
			a.h.Unregister <- c
			_ = c.Close()
		}()

		// Extract the JWT token from the query parameters
		tokenString := c.Locals("token").([]string)
		if tokenString == nil {
			log.Fatal("Missing token")
			return
		}

		// Validate the JWT token
		claims, err := midleware.ValidateJWT(tokenString[0], os.Getenv("JWT_SECRET"))
		if err != nil {
			log.Fatal("Invalid token:", err)
			return
		}

		// Authorize the user based on the token claims
		userID := claims["id"].(string)
		routeID := c.Query("route_id")

		a.h.Register <- c
		joinRoom(routeID, c)
		defer leaveRoom(routeID, c)

		for {
			var msg dto.Message
			err := c.ReadJSON(&msg)

			if err = helper.Validate.Struct(&msg); err != nil {
				log.Println(err)
			}

			if err != nil {
				log.Fatal(err)
				return
			}

			a.h.Broadcast <- dto.Message{
				UserID: userID,
				Lat:    msg.Lat,
				Lng:    msg.Lng,
			}

			_, err = a.repo.SaveCurrentDriverLocation(ctx, dto.Message{
				UserID: userID,
				Lat:    msg.Lat,
				Lng:    msg.Lng,
			})

			if err != nil {
				log.Println(err)
			}
		}
	}
}

func joinRoom(routeID string, ws *websocket.Conn) {
	roomsMu.Lock()
	room, exists := rooms[routeID]
	if !exists {
		room = &Room{clients: make(map[*websocket.Conn]bool)}
		rooms[routeID] = room
	}
	roomsMu.Unlock()

	room.mu.Lock()
	room.clients[ws] = true
	room.mu.Unlock()
}

func leaveRoom(routeID string, ws *websocket.Conn) {
	roomsMu.Lock()
	room, exists := rooms[routeID]
	if exists {
		room.mu.Lock()
		delete(room.clients, ws)
		room.mu.Unlock()
	}
	roomsMu.Unlock()
}

func NewWsGeoTracking(h *dto.Hub, repo repository.GeoTrackRepository) GeoTrack {
	return &GeoTrackImpl{
		h:    h,
		repo: repo,
	}
}
