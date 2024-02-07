package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	conn_type = "tcp"
	conn_addr = "localhost:8080"
	buff_size = 1024
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	server_conn, err := net.Dial(conn_type, conn_addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	go readFromServer(server_conn, os.Stdout)
	sendData(server_conn, reader)
}

type ReadFromConn interface {
	Read(b []byte) (n int, err error)
}

func readFromServer(conn net.Conn, writer io.Writer) {
	for {
		readConn(conn, writer)
	}
}

func readConn(conn net.Conn, writer io.Writer) {
	buffer := make([]byte, buff_size)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	buffer = bytes.Trim(buffer[:n], "\x00")
	fmt.Fprintln(writer, string(buffer))
}

type InputReader interface {
	ReadString(delim byte) (string, error)
}

func sendData(server_conn net.Conn, reader *bufio.Reader) {
	for {
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
