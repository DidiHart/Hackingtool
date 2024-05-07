package scanners

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func scanport(ctx context.Context, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	IP := "scanme.nmap.org"
	addr := fmt.Sprintf("%s:%d", IP, port)

	dialer := net.Dialer{}

	connection, err := dialer.DialContext(ctx, "tcp", addr)

	if err != nil {
		// fmt.Println("Connection error:", err)
		return
	}

	defer connection.Close()

	fmt.Printf("[+] Connection has been established... %v \n", connection.RemoteAddr().String())

}

func BasicScanner() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go scanport(ctx, i, &wg)
	}
	wg.Wait()
}
