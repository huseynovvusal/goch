package chat

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/huseynovvusal/goch/internal/config"
	"github.com/huseynovvusal/goch/internal/discovery"
)

type NetworkMessage struct {
	Content string
	From    discovery.NetworkUser
}

func SendChatMessage(content string, to discovery.NetworkUser, from discovery.NetworkUser) error {
	addr := &net.UDPAddr{
		IP:   net.ParseIP(to.IP),
		Port: config.CHAT_PORT,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	messageData := NetworkMessage{
		Content: content,
		From:    from,
	}
	data, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)

	return err
}

func ListenForChatMessages(messages chan<- []NetworkMessage) {
	addr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: config.CHAT_PORT,
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

		messages <- []NetworkMessage{message}
	}

}
