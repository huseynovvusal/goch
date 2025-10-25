package chat

import (
	"fmt"
	"net"
)

func StartServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Printf("Chat server started on port %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("New connection from", conn.RemoteAddr())
	defer conn.Close()
}
