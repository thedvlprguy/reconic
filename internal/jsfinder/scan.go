package jsfinder

import (
	"io"
	"net/http"
	"regexp"
	"time"
)

var jsRegex = regexp.MustCompile(`((http[s]?:\/\/)?[^\s"']+\.(js))`)
var endpointRegex = regexp.MustCompile(`(?i)(\/[a-zA-Z0-9_\-\/\.]*\?(?:[a-zA-Z0-9_\-]+=)?[a-zA-Z0-9_\-]*)`)

func ExtractJSLinks(htmlLinks []string) []string {
	var jsLinks []string
	for _, link := range htmlLinks {
		if jsRegex.MatchString(link) {
			jsLinks = append(jsLinks, link)
		}
	}
	return unique(jsLinks)
}

func ExtractEndpoints(jsURL string) []string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(jsURL)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}
	}

	matches := endpointRegex.FindAllString(string(body), -1)
	return unique(matches)
}

func unique(input []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, s := range input {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}
