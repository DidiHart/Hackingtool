package handler

import (
	"fmt"
	"net"
)

func ConnectWithServer(ServerIP string, Port string) (connection net.Conn, err error) {
	ServerAddress := fmt.Sprintf("%s:%s", ServerIP, Port)

	connection, err = net.Dial("tcp", ServerAddress)

	if err != nil {
		return
	}

	return
}
