package socket_io

import (
	"fmt"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/repository"
	"github.com/gofiber/contrib/socketio"
	"log"
)

type GeoTrack interface {
	JoinRoom(kws *socketio.Websocket)
	LeaveRoom(ep *socketio.EventPayload)
	Disconnect(ep *socketio.EventPayload)
	EventMessage(ep *socketio.EventPayload)
	SendLocation(ep *socketio.EventPayload)
}

type GeoTrackImpl struct {
	repo    repository.GeoTrackRepository
	clients map[string][]string
}

func (a *GeoTrackImpl) Disconnect(ep *socketio.EventPayload) {
	roomId := ep.Kws.Params("roomId")
	for i, v := range a.clients[roomId] {
		if v == ep.Kws.UUID {
			a.clients[roomId] = append(a.clients[roomId][:i], a.clients[roomId][i+1:]...)
			break
		}
	}
	log.Printf("Client %s disconnected from room: %s", ep.Kws.UUID, roomId)
}

func (a *GeoTrackImpl) SendLocation(ep *socketio.EventPayload) {
	roomId := ep.Kws.Params("roomId")
	ep.Kws.EmitToList(a.clients[roomId], []byte("message"), socketio.TextMessage)
}

func (a *GeoTrackImpl) JoinRoom(kws *socketio.Websocket) {
	roomId := kws.Params("roomId")
	a.clients[roomId] = append(a.clients[roomId], kws.UUID)
	log.Printf("Client %s joined room: %s", kws.UUID, roomId)
}

func (a *GeoTrackImpl) LeaveRoom(ep *socketio.EventPayload) {
	roomId := ep.Kws.Params("roomId")
	for i, v := range a.clients[roomId] {
		if v == ep.Kws.UUID {
			a.clients[roomId] = append(a.clients[roomId][:i], a.clients[roomId][i+1:]...)
			break
		}
	}
	log.Printf("Client %s leaved room: %s", ep.Kws.UUID, roomId)
}

func (a *GeoTrackImpl) EventMessage(ep *socketio.EventPayload) {
	fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))

	message := "test message"

	// Emit the message directly to specified user
	err := ep.Kws.EmitTo("1", []byte(message), socketio.TextMessage)
	if err != nil {
		fmt.Println(err)
	}
}

//func (a *GeoTrackImpl) Run() {
//	for {
//		select {
//		case conn := <-a.h.Register:
//			if _, ok := a.h.Rooms[conn.ID]; !ok {
//				a.h.Clients[conn] = true
//			}
//		case conn := <-a.h.Unregister:
//			delete(a.h.Clients, conn)
//		case conn := <-a.h.Rooms:
//
//		case msg := <-a.h.Broadcast:
//			for conn := range a.h.Clients {
//				_ = conn.WriteJSON(msg)
//			}
//		}
//	}
//}
//
//func (a *GeoTrackImpl) LocationTracking(ctx context.Context) func(*websocket.Conn) {
//	return func(c *websocket.Conn) {
//		routeID := c.Params("route_id")
//
//		defer func() {
//			a.h.Unregister <- c
//			_ = c.Close()
//		}()
//
//		a.h.Register <- c
//		joinRoom(routeID, c)
//		defer leaveRoom(routeID, c)
//
//		for {
//			select {
//			case <-ctx.Done():
//				log.Println("Connection closed")
//				return
//			default:
//			}
//
//			var msg dto.Message
//			err := c.ReadJSON(&msg)
//
//			if err != nil {
//				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
//					log.Println("WebSocket connection closed:", err)
//					return
//				}
//				log.Println("Error reading JSON:", err)
//				return
//			}
//
//			if err = helper.Validate.Struct(&msg); err != nil {
//				log.Println("Validation error:", err)
//				_ = c.WriteJSON(map[string]string{"error": "Invalid message format"})
//				continue
//			}
//
//			a.h.Broadcast <- dto.Message{
//				UserID: userID,
//				Lat:    msg.Lat,
//				Lng:    msg.Lng,
//			}
//
//			_, err = a.repo.SaveCurrentDriverLocation(ctx, dto.Message{
//				UserID: userID,
//				Lat:    msg.Lat,
//				Lng:    msg.Lng,
//			})
//
//			if err != nil {
//				log.Println("Error saving driver location:", err)
//				_ = c.WriteJSON(map[string]string{"error": "Could not save location"})
//				continue
//			}
//		}
//	}
//}
//
//func (h *Handler) JoinRoom(c *fiber.Ctx) func(*websocket.Conn) {
//	return func(c *websocket.Conn) {
//		roomID := c.Params("roomID")
//
//		tokenString := c.Locals("token").([]string)
//		if tokenString == nil {
//			log.Fatal("Missing token")
//			return
//		}
//
//		// Validate the JWT token
//		claims, err := midleware.ValidateJWT(tokenString[0], os.Getenv("JWT_SECRET"))
//		if err != nil {
//			log.Fatal("Invalid token:", err)
//			return
//		}
//
//		// Authorize the user based on the token claims
//		userID := claims["id"].(string)
//
//		client := &dto.Client{
//			Conn:     c,
//			Username: username,
//			RoomID:   roomID,
//		}
//
//		h.hub.Register <- client
//
//	}
//}

func NewWsGeoTracking(repo repository.GeoTrackRepository) GeoTrack {
	return &GeoTrackImpl{
		repo: repo,
	}
}
