package hub

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/mikroNet/notification/config/rabbitmq"
	"github.com/GabrielMoody/mikroNet/notification/internal/dto"
)

type Hub struct {
	clients map[string]net.Conn
	amqp    *common.AMQP
	mu      sync.Mutex
}

func (h *Hub) SendOrderNotification() {
	msg, err := h.amqp.Consume("order_notifications", "order", "order.notification")

	if err != nil {
		log.Fatal(err.Error())
	}

	var req dto.OrderNotificationData

	go func() {
		for m := range msg {
			if err := json.Unmarshal(m.Body, &req); err != nil {
				log.Fatal(err)
			}

			h.mu.Lock()

			conn, ok := h.clients[req.RecipientID]

			if !ok {
				log.Fatal("client id doesn't exist")
			}

			b, _ := json.Marshal(req)

			conn.Write(b)

			h.mu.Unlock()
		}
	}()

	select {}
}

func (n *Hub) HandleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var userID string

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", userID)
			n.Unregister(userID)
			return
		}

		msg = strings.TrimSpace(msg)

		if strings.HasPrefix(msg, "IDENTIFY ") {
			// IDENTIFY <userID>
			userID = strings.TrimPrefix(msg, "IDENTIFY ")
			n.Register(userID, conn)
			fmt.Println("Registered:", userID)
			continue
		}

		fmt.Println("Message from", userID, "->", msg)
	}
}

func (n *Hub) Register(id string, conn net.Conn) {
	n.mu.Lock()
	n.clients[id] = conn
	n.mu.Unlock()
}

func (n *Hub) Unregister(id string) {
	if id == "" {
		return
	}

	n.mu.Lock()
	delete(n.clients, id)
	n.mu.Unlock()
}

func NewHub() *Hub {
	amqp := rabbitmq.Init("amqp://admin:admin123@localhost:5672/")
	return &Hub{
		clients: make(map[string]net.Conn),
		amqp:    amqp,
	}
}
