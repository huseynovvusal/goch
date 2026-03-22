package tui

import (
	"github.com/huseynovvusal/goch/internal/chat"
	"github.com/huseynovvusal/goch/internal/discovery"
)

func addDummyData(m *Model) {
	dummies := []discovery.NetworkUser{
		{Name: "Alice_go", IP: "192.168.1.101"},
		{Name: "Bob_builder", IP: "192.168.1.102"},
		{Name: "Cyber_ninja", IP: "192.168.1.103"},
	}

	for _, dummy := range dummies {
		found := false
		for _, u := range m.onlineUsers {
			if u.Name == dummy.Name {
				found = true
				break
			}
		}
		if !found {
			m.onlineUsers = append(m.onlineUsers, dummy)
		}
	}

	if len(m.chatMessages) == 0 {
		m.chatMessages = []chat.NetworkMessage{
			{From: discovery.NetworkUser{Name: "Alice_go", IP: "192.168.1.101"}, Content: "Hey, nice Earthy theme!"},
			{From: discovery.NetworkUser{Name: m.username, IP: ""}, Content: "Thanks! It looks very G.O.A.T."},
		}
	}
}
