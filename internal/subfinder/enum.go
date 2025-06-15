package subfinder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var apis = []func(string) []string{
	fromCrtsh,
	fromHackertarget,
	fromAlienvault,
}

func Enumerate(domain string) []string {
	var all []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, api := range apis {
		wg.Add(1)
		go func(apiFunc func(string) []string) {
			defer wg.Done()
			subs := apiFunc(domain)
			mu.Lock()
			all = append(all, subs...)
			mu.Unlock()
		}(api)
	}

	wg.Wait()
	return dedupe(all)
}

func fromCrtsh(domain string) []string {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data []map[string]interface{}
	json.Unmarshal(body, &data)

	var subs []string
	for _, entry := range data {
		if name, ok := entry["name_value"].(string); ok {
			parts := strings.Split(name, "\n")
			subs = append(subs, parts...)
		}
	}
	return subs
}

func fromHackertarget(domain string) []string {
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\n")

	var subs []string
	for _, line := range lines {
		if fields := strings.Split(line, ","); len(fields) > 0 {
			subs = append(subs, fields[0])
		}
	}
	return subs
}

func fromAlienvault(domain string) []string {
	url := fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/passive_dns", domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	var subs []string
	if passiveDNS, ok := result["passive_dns"].([]interface{}); ok {
		for _, item := range passiveDNS {
			if record, ok := item.(map[string]interface{}); ok {
				if hostname, ok := record["hostname"].(string); ok {
					subs = append(subs, hostname)
				}
			}
		}
	}
	return subs
}

func dedupe(list []string) []string {
	seen := make(map[string]bool)
	var unique []string
	for _, item := range list {
		item = strings.TrimSpace(item)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		unique = append(unique, item)
	}
	return unique
}
