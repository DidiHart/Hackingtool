package macchanger

import (
	"flag"

	"github.com/didihart/hacking_tool/commands"
)

func MacChanger(command string, args []string) {

	intFace := flag.String("interface", "", "Interface for mac changing")
	macAddr := flag.String("macaddr", "", "Input new MAC address")

	flag.Parse()

	commands.ExecuteCommands("sudo", []string{"ifconfig", *intFace, "down"})
	commands.ExecuteCommands("sudo", []string{"ifconfig", *intFace, "hw", "ether", *macAddr})
	commands.ExecuteCommands("sudo", []string{"ifconfig", *intFace, "up"})
}
