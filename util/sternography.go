package util

import (
	"fmt"
	"os"

	"github.com/DimitarPetrov/stegify/steg"
)

func StegHide(carrierFile, malFile, encFile string) {
	err := steg.EncodeByFileNames(carrierFile, malFile, encFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("file has been encoded succesfully")
}
