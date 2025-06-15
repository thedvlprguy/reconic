package resolver

import (
	"context"
	"net"
	"sync"
	"time"
)

func ResolveSubdomains(subdomains []string) []string {
	var live []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: 2 * time.Second}
			return d.DialContext(ctx, "udp", "8.8.8.8:53") // Use Google's DNS
		},
	}

	for _, sub := range subdomains {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if _, err := resolver.LookupHost(ctx, s); err == nil {
				mu.Lock()
				live = append(live, s)
				mu.Unlock()
			}
		}(sub)
	}

	wg.Wait()
	return live
}
