package server

import (
	"net"

	"github.com/rickychandra/goboil/common/logger"
)

//CreateTCPListener creates a TCP listener for host and port. Port is in the form ":xxxx".
//If host is empty string, will listen on all interfaces.
var CreateTCPListener = func(host, port string) net.Listener {
	lis, err := net.Listen("tcp", host+port)
	if err != nil {
		logger.Error(nil, err)
	}
	return lis
}
