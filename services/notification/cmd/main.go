package main

import (
	"fmt"
	"net"

	"github.com/GabrielMoody/mikroNet/notification/internal/hub"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	hub := hub.NewHub()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())
		go hub.HandleClient(conn)

		hub.SendOrderNotification()
	}
}
