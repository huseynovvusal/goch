package chat

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/huseynovvusal/goch/internal/discovery"
)

type NetworkMessage struct {
	Content   string
	From      discovery.NetworkUser
	Timestamp time.Time `json:"timestamp"`
}

func SendChatMessage(content string, to discovery.NetworkUser, from discovery.NetworkUser, chatPort int) error {
	addr := &net.UDPAddr{
		IP:   net.ParseIP(to.IP),
		Port: chatPort,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	messageData := NetworkMessage{
		Content:   content,
		From:      from,
		Timestamp: time.Now(),
	}
	data, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)

	return err
}

func ListenForChatMessages(messages chan<- NetworkMessage, chatPort int) {
	addr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: chatPort,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening for chat messages:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		var message NetworkMessage
		if err := json.Unmarshal(buf[:n], &message); err != nil {
			continue
		}

		messages <- message
	}

}
