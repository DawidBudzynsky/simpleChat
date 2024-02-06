package main

import (
	"fmt"
	"net"
)

func main() {
	server_sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		client_sock, err := server_sock.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(client_sock)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received from client: %s\n", buf)
}
