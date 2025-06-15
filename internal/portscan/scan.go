package portscan

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

var commonPorts = []int{21, 22, 23, 25, 53, 80, 110, 111, 135, 139, 143, 443, 445, 993, 995, 1723, 3306, 3389, 5900, 8080, 8443}

func Scan(host string, ports []int) []int {
	var open []int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, port := range ports {
		port := port // capture loop variable correctly
		wg.Add(1)
		go func() {
			defer wg.Done()
			address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err == nil {
				conn.Close()
				mu.Lock()
				open = append(open, port)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	sort.Ints(open)
	return open
}

func CommonPorts() []int {
	return commonPorts
}
