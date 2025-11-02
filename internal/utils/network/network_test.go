package network

import (
	"net"
	"testing"
)

func TestGetBroadcastAddr(t *testing.T) {
	port := 9999
	addr, err := GetBroadcastAddr(port)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if addr == nil {
		t.Fatal("Expected non-nil address")
	}
	if addr.Port != port {
		t.Errorf("Expected port %d, got %d", port, addr.Port)
	}
	ip := addr.IP.To4()
	if ip == nil {
		t.Errorf("Expected an IPv4 address, got %v", addr.IP)
	}
	if ip.IsLoopback() {
		t.Errorf("Expected non-loopback address, got %v", ip)
	}
	if ip.IsUnspecified() {
		t.Errorf("Expected specified address, got %v", ip)
	}

	ifaces, _ := net.Interfaces()
	found := false
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue
			}
			mask := ipnet.Mask
			localIP := ipnet.IP.To4()
			broadcast := make(net.IP, 4)
			for i := 0; i < 4; i++ {
				broadcast[i] = localIP[i] | ^mask[i]
			}
			if broadcast.Equal(ip) {
				found = true
				break
			}
		}
	}
	if !found {
		t.Errorf("Returned IP %v is not a broadcast address for any active interface", ip)
	}
}
