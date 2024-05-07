package commands

import (
	"log"
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
