package handler

import (
	"fmt"
	"net"
)

func ConnectWithVictim(IP string, port string) (connection net.Conn, err error) {
	LocalAddressPort := fmt.Sprintf("%s:%s", IP, port)

	listener, err := net.Listen("tcp", LocalAddressPort)
	if err != nil {
		return
	}
	connection, err = listener.Accept()
	if err != nil {
		return
	}
	return
}
