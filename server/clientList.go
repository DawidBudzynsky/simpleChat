package main

import (
	"net"
	"sync"
)

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
