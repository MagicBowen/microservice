package registration

import (
	"errors"
	"log"
	"net"
)

const localhost = "localhost"

func getNoneLoopbackIpv4(addr net.Addr) (string, error) {
	ipnet, ok := addr.(*net.IPNet)
	if !ok {
		return "", errors.New("get ip error")
	}
	if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
		return ipnet.IP.String(), nil
	}
	return "", errors.New("ip not satisfied")
}

func getLocalIP() string {
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, addr := range addrs {
			if ip, err := getNoneLoopbackIpv4(addr); err == nil {
				return ip
			}
		}
	}
	log.Printf("get none satisfied IP address")
	return localhost
}
