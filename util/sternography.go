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

func stegUnMask(encodedFile, resultFile string) {
	err := steg.DecodeByFileNames(encodedFile, resultFile)

	if err != nil {
		fmt.Println("can't decode file", err)
		os.Exit(1)
	}
	fmt.Println("decoded file succesfully")
}
