package chat_test

import (
	"encoding/json"
	"net"
	"testing"
	"time"

	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/db"
	"github.com/huseynovvusal/goch/internal/discovery"
)

func TestSendChatMessage(t *testing.T) {
	// This function is a placeholder for testing SendChatMessage.
	// In a real test, you would set up a mock UDP server to receive the message
	// and verify that the content and sender information are correct.
	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: db.DefaultChatPort,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("Failed to set up UDP listener: %v", err)
	}
	defer conn.Close()

	from := discovery.NetworkUser{IP: "127.0.0.1", Name: "Sender"}
	to := discovery.NetworkUser{IP: "127.0.0.1", Name: "Receiver"}
	content := "Hello, World!"

	errCh := make(chan error, 1)
	go func() {
		errCh <- chat.SendChatMessage(content, to, from, db.DefaultChatPort)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("Failed to send chat message: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Timed out sending chat message")
	}

	buf := make([]byte, 1024)
	if err := conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil { // Don't wait forever
		t.Fatalf("Failed to set read deadline: %v", err)
	}
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		t.Fatalf("Failed to read sent message: %v", err)
	}

	var received chat.NetworkMessage
	err = json.Unmarshal(buf[:n], &received)
	if err != nil {
		t.Fatalf("Failed to unmarshal received data: %v", err)
	}

	if received.Content != content {
		t.Errorf("Expected content %s, got %s", content, received.Content)
	}
}

func TestListenForChatMessages(t *testing.T) {
	messages := make(chan chat.NetworkMessage, 1)

	go chat.ListenForChatMessages(messages, db.DefaultChatPort)

	srvrAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: db.DefaultChatPort,
	}

	conn, err := net.DialUDP("udp", nil, srvrAddr)
	if err != nil {
		t.Fatalf("Failed to set up UDP connection: %v", err)
	}
	defer conn.Close()

	testMessage := chat.NetworkMessage{
		Content:   "Test Message",
		From:      discovery.NetworkUser{IP: "127.0.0.1", Name: "Sender"},
		Timestamp: time.Now(),
	}
	data, err := json.Marshal(testMessage)
	if err != nil {
		t.Fatalf("Failed to marshal test message: %v", err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatalf("Failed to write test message: %v", err)
	}

}
