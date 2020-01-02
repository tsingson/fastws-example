package vtils

import (
	"net"

	"emperror.dev/errors"
)

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, errors.New("can't get ip")
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	_ = conn.Close()

	return localAddr.IP, nil
}
