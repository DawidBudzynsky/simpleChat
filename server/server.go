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
		buf, ok := ReadFromClient(senderConn, clientList)
		if !ok {
			return
		}

		for client := range clientList.clients {
			if client != senderConn {
				SendToClient(client, buf)
			}
		}
	}
}

func ReadFromClient(client net.Conn, clientList *ClientList) (buf []byte, ok bool) {
	buf = make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		clientList.Remove(client)
		fmt.Println(clientList.List())
		return nil, false
	}
	// NOTE: removing trailing zeros prob bcs of 1024 ^
	buf = buf[:n]
	return buf, true
}

func SendToClient(client net.Conn, buf []byte) {
	_, err := client.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
}
