package discovery

import (
	"fmt"
	"net"
	"time"

	"github.com/huseynovvusal/goch/internal/utils/network"
)

type NetworkUser struct {
	Name string
	IP   string
}

var onlineUsers = []NetworkUser{}

func GetOnlineUsers() []NetworkUser {
	return onlineUsers
}

func BroadcastPresence(name string, port int) {
	addr, err := network.GetBroadcastAddr(port)
	if err != nil {
		fmt.Println("Error getting broadcast address:", err)
		return
	}

	conn, _ := net.DialUDP("udp", nil, addr)
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
		user := NetworkUser{Name: name, IP: remoteAddr.IP.String()}

		if !isUserInList(user) {
			onlineUsers = append(onlineUsers, user)
		}

		// fmt.Printf("Discovered user: %s at %s; online users: %v\n", user.Name, user.IP, onlineUsers)
	}
}

func isUserInList(user NetworkUser) bool {
	for _, u := range onlineUsers {
		if u.IP == user.IP {
			return true
		}
	}

	return false
}
