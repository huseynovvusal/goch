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

var onlineUsers = []User{}

func GetOnlineUsers() []User {
	return onlineUsers
}

func BroadcastPresence(name string, port int) {
	addr := net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: port,
	}

	conn, _ := net.DialUDP("udp", nil, &addr)
	defer conn.Close()

	for {
		conn.Write([]byte(name))
		// fmt.Println("Broadcasted presence as", name)
		time.Sleep(3 * time.Second)
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

		if !isUserInList(user) {
			onlineUsers = append(onlineUsers, user)
		}

		// fmt.Printf("Discovered user: %s at %s; online users: %v\n", user.Name, user.IP, onlineUsers)
	}
}

func isUserInList(user User) bool {
	for _, u := range onlineUsers {
		if u.IP == user.IP {
			return true
		}
	}

	return false
}
