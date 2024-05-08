package filenavigation

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NavigateSystem(conn net.Conn) (err error) {

	ConnectionReader := bufio.NewReader(conn)

	initialPwdRaw, err := ConnectionReader.ReadString('\n')

	initialPwd := strings.TrimSuffix(initialPwdRaw, "\n")

	looper := true

	for looper {
		fmt.Print(initialPwd, " >> ")

		CommandReader := bufio.NewReader(os.Stdin)

		userCommandRaw, err := CommandReader.ReadString('\n')

		if err != nil {
			fmt.Println("[+] Unable to read command ")
		}

		_, err = conn.Write([]byte(userCommandRaw))

		if err != nil {
			fmt.Println("[+] conn error ", err)
		}

		if userCommandRaw == "stop\n" {
			looper = false
			break
		}

		newPwd, _ := ConnectionReader.ReadString('\n')

		fmt.Println("[+] Working Directory changed to ", newPwd)

		initialPwd = newPwd
	}
	return
}
