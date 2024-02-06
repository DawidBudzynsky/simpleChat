package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	server_conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server_conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(conn net.Conn, wg *sync.WaitGroup) {
		defer wg.Done()
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from server: %s\n", buffer)
	}(server_conn, &wg)

	_, err = server_conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.Wait()
}
