package handler

import (
	"fmt"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/socket_io"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func NewWSHandler(r fiber.Router, db *gorm.DB) {
	api := r.Group("/")

	repo := repository.NewGeoTrackRepository(db)
	socket := socket_io.NewWsGeoTracking(repo)

	//socketio.On(socketio.EventConnect, socket.JoinRoom)
	socketio.On(socketio.EventClose, socket.LeaveRoom)
	socketio.On(socketio.EventDisconnect, socket.Disconnect)
	socketio.On(socketio.EventMessage, socket.EventMessage)
	socketio.On("location", socket.SendLocation)
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		fmt.Printf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	api.Get("/ws/:roomId", socketio.New(socket.JoinRoom))
}
