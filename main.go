package main

import "github.com/didihart/hacking_tool/util"

func main() {
	// scanners.BasicScanner()
	// util.NetworkScanner()
	// util.PacketSniffer()
	carrierFile := "test.png"
	// maliciousFile := "helix_bane.bat"
	// maliciousFile := "helix_bane.ps1"
	maliciousFile := "norvoinc.exe"
	encodedFile := "helix_test.png"

	util.StegHide(carrierFile, maliciousFile, encodedFile)
}
