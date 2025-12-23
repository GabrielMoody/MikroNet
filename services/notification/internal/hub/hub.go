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
	"github.com/GabrielMoody/MikroNet/services/notification/config/rabbitmq"
	"github.com/GabrielMoody/MikroNet/services/notification/internal/dto"
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

	go func() {
		for m := range msg {

			var req dto.OrderNotificationData

			if err := json.Unmarshal(m.Body, &req); err != nil {
				log.Fatal(err)
				m.Nack(false, false)
				continue
			}

			h.mu.Lock()

			conn, ok := h.clients[req.RecipientID]

			h.mu.Unlock()

			if !ok {
				log.Fatal("client id doesn't exist")
				m.Nack(false, false)
				continue
			}

			b, _ := json.Marshal(req)

			if _, err = conn.Write(append(b, '\n')); err != nil {
				log.Printf("client %s disconnected, removing\n", req.RecipientID)

				h.mu.Lock()
				conn.Close()
				delete(h.clients, req.RecipientID)
				h.mu.Unlock()

				m.Nack(false, false)
				continue
			}

			m.Ack(false)
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
