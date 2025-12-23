package main

import (
	"fmt"
	"net"

	"github.com/GabrielMoody/MikroNet/services/notification/internal/hub"
)

func main() {
	hub := hub.NewHub()
	go hub.SendOrderNotification()

	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())
		go hub.HandleClient(conn)
	}
}
