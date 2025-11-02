package network

import (
	"errors"
	"net"
)

func GetBroadcastAddr(port int) (*net.UDPAddr, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue // interface down or loopback
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue // not an ipv4 address
			}

			ip := ipnet.IP.To4()
			mask := ipnet.Mask
			broadcast := make(net.IP, 4)

			for i := 0; i < 4; i++ {
				broadcast[i] = ip[i] | ^mask[i]
			}

			return &net.UDPAddr{
				IP:   broadcast,
				Port: port,
			}, nil
		}
	}

	return nil, errors.New("no suitable network interface found")
}
