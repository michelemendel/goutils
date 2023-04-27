package network

import (
	"net"

	"github.com/michelemendel/goutils/log"
)

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Errorf("Couldn't get an IP address, %v", err)
		return ""
	}
	for _, address := range addrs {
		// Check the address type and if it is not a loopback then display it
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	log.Errorf("Couldn't get an IP address, %v", err)
	return ""
}
