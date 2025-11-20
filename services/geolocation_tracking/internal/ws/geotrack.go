package ws

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/helper"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/midleware"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/gofiber/contrib/websocket"
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
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Closing"))
			_ = c.Close()
		}()

		tokenString := c.Query("token")
		var userID string

		if tokenString != "" {
			claims, err := midleware.ValidateJWT(tokenString, os.Getenv("JWT_SECRET"))
			if err != nil {
				log.Println("Invalid token:", err)
				// Optionally: return here if you want to reject invalid tokens
			} else {
				if id, ok := claims["id"].(string); ok {
					userID = id
				}
			}
		} else {
			log.Println("No token provided, proceeding without user ID")
		}

		a.h.Register <- c

		c.SetPongHandler(func(appData string) error {
			log.Println("Received pong")
			return nil
		})

		go func() {
			for {
				time.Sleep(30 * time.Second)
				if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Println("Ping error:", err)
					return
				}
			}
		}()

		for {
			var msg dto.Message
			err := c.ReadJSON(&msg)

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("Unexpected WebSocket close:", err)
				}
				return
			}

			if err = helper.Validate.Struct(&msg); err != nil {
				log.Println(err)
			}

			if err != nil {
				log.Println(err)
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

func NewWsGeoTracking(h *dto.Hub, repo repository.GeoTrackRepository) GeoTrack {
	return &GeoTrackImpl{
		h:    h,
		repo: repo,
	}
}
