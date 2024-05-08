package commands

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func ExecuteCommands(command string, args []string) (err error) {
	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stderr = os.Stderr
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stdin = os.Stdin

	err = cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
		return
	}

	return nil
}

type Command struct {
	CmdOutput string
	CmdError  string
}

// send command from server
func ExecuteCommandWindows(conn net.Conn) (err error) {

	reader := bufio.NewReader(os.Stdin)

	looper := true

	for looper {

		fmt.Print(">> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		conn.Write([]byte(command))
		if command == "stop\n" { //condition to stop
			looper = false
			continue
		}

		cmdStruct := &Command{}

		decoder := gob.NewDecoder(conn)
		err = decoder.Decode(cmdStruct)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(cmdStruct.CmdOutput)

		if cmdStruct.CmdError != "" {
			fmt.Println(cmdStruct.CmdError)
		}
	}
	return
}
