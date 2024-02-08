package main

import (
	"bytes"
	"net"
	"testing"
)

func TestReadFromClient(t *testing.T) {
	t.Run("Reading data from client", func(t *testing.T) {
		clientList := NewClientList()
		clientConn, serverConn := net.Pipe()
		clientList.Add(clientConn)

		go func() {
			text := "Hello"
			serverConn.Write([]byte(text))
		}()

		buf, ok := ReadFromClient(clientConn, clientList)

		expectedBuf := []byte("Hello")
		expectedOk := true
		if ok != expectedOk {
			t.Errorf("Expected ok: %t, got: %t", expectedOk, ok)
		}
		if len(buf) != len(expectedBuf) {
			t.Errorf("Expected buf length: %d, got: %d", len(expectedBuf), len(buf))
		}
		if !bytes.Equal(buf, expectedBuf) {
			t.Errorf("Different message: %s, got: %s", expectedBuf, buf)
		}
		clientConn.Close()
	})
}
