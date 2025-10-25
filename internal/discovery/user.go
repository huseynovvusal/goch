package discovery

import (
	"fmt"
	"net"
	"time"
)

type User struct {
	Name string
	IP   string
}

var onlineUsers = []User{
	// Placeholder for demo
	{Name: "John Doe", IP: "192.168.10.10"},
}

func GetOnlineUsers() []User {
	return onlineUsers
}

func BroadcastPresence(name string, port int) {
	addr := net.UDPAddr{
		IP:   net.ParseIP("192.168.100.255"),
		Port: port,
	}

	conn, _ := net.DialUDP("udp", nil, &addr)
	defer conn.Close()

	for {
		conn.Write([]byte(name))
		fmt.Println("Broadcasted presence as", name)
		time.Sleep(5 * time.Second)
	}
}

func ListenForPresence(port int) {
	addr := net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening for presence:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		name := string(buf[:n])
		user := User{Name: name, IP: remoteAddr.IP.String()}
		onlineUsers = append(onlineUsers, user)
	}
}
