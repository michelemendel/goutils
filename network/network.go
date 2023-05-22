package network

import (
	"net"

	"github.com/michelemendel/goutils/log"
	"go.uber.org/zap"
)

var lg *zap.SugaredLogger

const LOG_LEVEL = "INFO"

func init() {
	lg = log.InitWithConsole(LOG_LEVEL)
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		lg.Errorf("Couldn't get an IP address, %v", err)
		return ""
	}
	for _, address := range addrs {
		// Check the address type and if it is not a loopback then display it
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	lg.Errorf("Couldn't get an IP address, %v", err)
	return ""
}
