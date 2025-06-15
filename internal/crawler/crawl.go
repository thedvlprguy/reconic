package crawler

import (
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func Crawl(url string) []string {
	var links []string

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := client.Get(url)
	if err != nil {
		return links
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		t := z.Token()
		if t.Type == html.StartTagToken {
			for _, attr := range t.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					link := attr.Val
					if strings.HasPrefix(link, "/") {
						link = url + link
					}
					if strings.HasPrefix(link, "http") {
						links = append(links, link)
					}
				}
			}
		}
	}

	return unique(links)
}

func unique(input []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, link := range input {
		if !seen[link] {
			seen[link] = true
			result = append(result, link)
		}
	}
	return result
}
