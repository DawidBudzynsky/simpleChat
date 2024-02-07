package main

import (
	"net"
	"strings"
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

type MockReader struct {
	word string
}

func (r *MockReader) ReadString(delim byte) (string, error) {
	return r.word, nil
}

func TestSendData(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	reader := &MockReader{word: "Im sending this"}

	go func() {
		sendData(serverConn, reader)
	}()
	output := ReadSentData(t, clientConn)

	expected := "Im sending this"
	if output != expected {
		t.Errorf("Expected output: %q, got: %q", expected, output)
	}
}

func ReadSentData(t testing.TB, clientConn net.Conn) (output string) {
	t.Helper()

	buffer := make([]byte, 1024)
	n, err := clientConn.Read(buffer)
	if err != nil {
		t.Fatalf("Error reading from client connection: %v", err)
	}
	return strings.TrimSpace(string(buffer[:n]))
}
