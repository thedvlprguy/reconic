package secrets

import (
	"io"
	"net/http"
	"regexp"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(api[_-]?key\s*=\s*['"][A-Za-z0-9_\-]{16,}['"])`),
	regexp.MustCompile(`(?i)(access[_-]?token\s*=\s*['"][A-Za-z0-9\-_.]{10,}['"])`),
	regexp.MustCompile(`(?i)(authorization\s*[:=]\s*['"][A-Za-z0-9\-_.:]{10,}['"])`),
	regexp.MustCompile(`(?i)(firebase[_-]config\s*=\s*\{[^}]+\})`),
	regexp.MustCompile(`(?i)(aws[_-]secret[_-]access[_-]?key\s*=\s*['"][A-Za-z0-9/+=]{40}['"])`),
}

func FindSecrets(jsURL string) []string {
	resp, err := http.Get(jsURL)
	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var found []string
	content := string(body)

	for _, pattern := range patterns {
		matches := pattern.FindAllString(content, -1)
		for _, match := range matches {
			found = append(found, match)
		}
	}

	return found
}
