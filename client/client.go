package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	server_conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	go readFromServer(server_conn)
	sendData(server_conn)
}

func readFromServer(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from server: %s\n", buffer)
	}
}

func sendData(server_conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input: ", err)
			continue
		}
		text = strings.TrimSpace(text)

		_, err = server_conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
