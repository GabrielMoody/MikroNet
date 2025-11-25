package handler

import (
	"context"

	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewWSHandler(r fiber.Router, db *gorm.DB, rdb *redis.Client) {
	api := r.Group("/")

	ctx := context.Background()
	repo := repository.NewGeoTrackRepository(db, rdb)
	hub := &dto.Hub{
		Broadcast:  make(chan dto.Message),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Clients:    make(map[*websocket.Conn]bool),
	}
	wsLocation := ws.NewWsGeoTracking(hub, repo)

	go wsLocation.Run()

	api.Use("/ws/location", websocket.New(wsLocation.LocationTracking(ctx)))

}
