package ws

import (
	"context"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/helper"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/gofiber/contrib/websocket"
	"log"
)

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

		a.h.Register <- c

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

			a.h.Broadcast <- msg

			_, err = a.repo.SaveCurrentDriverLocation(ctx, msg)

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
