package net

import (
	"flag"
	"fmt"
	"sync"
)
import "net"

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func ScanNet() error {
	flag.Parse()
	var wg sync.WaitGroup
	// Escanear cada puerto y hacer una conexión
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("site:", *site, "Port", port, "is open")

		}(i)
	}
	wg.Wait()
	return nil
}
