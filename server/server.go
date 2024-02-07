package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	clientList := NewClientList()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		client_sock, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		clientList.Add(client_sock)
		fmt.Println(clientList.List())
		go handleConnection(client_sock, clientList)
	}
}

func handleConnection(senderConn net.Conn, clientList *ClientList) {
	for {
		buf := make([]byte, 1024)
		_, err := senderConn.Read(buf)
		if err != nil {
			clientList.Remove(senderConn)
			fmt.Println(clientList.List())
			return
		}

		for client := range clientList.clients {
			if client != senderConn {
				_, err = client.Write(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
