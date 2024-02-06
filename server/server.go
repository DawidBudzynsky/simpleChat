package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	clientList := NewClientList()
	ln, err := net.Listen("tcp", ":8080")
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

type ClientList struct {
	clients map[net.Conn]struct{}
	sync.RWMutex
}

func NewClientList() *ClientList {
	return &ClientList{
		clients: make(map[net.Conn]struct{}),
	}
}

func (cl *ClientList) Add(client net.Conn) {
	cl.Lock()
	defer cl.Unlock()
	cl.clients[client] = struct{}{}
}

func (cl *ClientList) Remove(client net.Conn) {
	cl.Lock()
	defer cl.Unlock()
	delete(cl.clients, client)
}

func (cl *ClientList) List() []net.Conn {
	cl.RLock()
	defer cl.RUnlock()
	clients := make([]net.Conn, 0, len(cl.clients))
	for client := range cl.clients {
		clients = append(clients, client)
	}
	return clients
}

func handleConnection(senderConn net.Conn, clientList *ClientList) {
	for {
		buf := make([]byte, 1024)
		_, err := senderConn.Read(buf)
		if err != nil {
			fmt.Println(err)
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
		fmt.Printf("Received from client: %s\n", buf)
	}
}
