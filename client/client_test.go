package main

import (
	"net"
	"sync"
	"testing"
)

type MockWriter struct {
	output string
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	w.output += string(p)
	return len(p), nil
}

func TestReadFromServer(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	writer := &MockWriter{}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		readConn(serverConn, writer)
	}()

	// simulate server
	go func() {
		text := string("Hello!")
		_, err := clientConn.Write([]byte(text))
		if err != nil {
			t.Errorf("Error writing to server: %v", err)
		}
	}()
	wg.Wait()

	output := writer.output

	expected := "Hello!\n"
	if output != expected {
		t.Errorf("Expected output: %q, got: %q", expected, output)
	}
}
