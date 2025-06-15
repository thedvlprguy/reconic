package main

import (
	"flag"
	"os"
	"strings"

	"github.com/pterm/pterm"

	"github.com/thedvlprguy/reconic/internal/crawler"
	"github.com/thedvlprguy/reconic/internal/jsfinder"
	"github.com/thedvlprguy/reconic/internal/resolver"
	"github.com/thedvlprguy/reconic/internal/secrets"
	"github.com/thedvlprguy/reconic/internal/subfinder"
)

func main() {
	var output string
	flag.StringVar(&output, "o", "", "Output file to save the subdomains")
	flag.Parse()

	if len(flag.Args()) < 1 {
		pterm.Error.Println("Usage: reconic <domain> [-o output.txt]")
		return
	}
	domain := flag.Args()[0]

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println(" Reconic - High-Speed Recon CLI ")
	pterm.DefaultBox.WithTitle("🎯 Target Domain").WithTitleTopLeft().Println(domain)

	spinner, _ := pterm.DefaultSpinner.Start("🔍 Gathering subdomains...")
	results := subfinder.Enumerate(domain)
	spinner.Stop()
	pterm.Success.Println("✅ Subdomain enumeration complete.")

	pterm.DefaultBox.WithTitle("🌐 Discovered Subdomains").WithTitleTopLeft().Println(strings.Join(results, "\n"))
	pterm.Info.Printf("🔢 Total Unique Subdomains: %d\n", len(results))

	if output != "" {
		err := os.WriteFile(output, []byte(strings.Join(results, "\n")), 0644)
		if err != nil {
			pterm.Error.Println("❌ Failed to save file:", err)
			return
		}
		pterm.Success.Printf("📁 Results saved to: %s\n", output)
	}

	pterm.Info.Println("🌐 Resolving live subdomains...")
	live := resolver.ResolveSubdomains(results)

	pterm.DefaultBox.WithTitle("🟢 Live Subdomains").WithTitleTopLeft().Println(strings.Join(live, "\n"))
	pterm.Info.Printf("🔢 Total Live Subdomains: %d\n", len(live))

	if output != "" {
		liveFile := strings.ReplaceAll(output, ".txt", "_live.txt")
		err := os.WriteFile(liveFile, []byte(strings.Join(live, "\n")), 0644)
		if err != nil {
			pterm.Error.Println("❌ Failed to save live file:", err)
			return
		}
		pterm.Success.Printf("📁 Live subdomains saved to: %s\n", liveFile)
	}

	var allLinks []string
	for _, liveDomain := range live {
		title := "🔎 Crawling: " + liveDomain
		links := crawler.Crawl(liveDomain)
		allLinks = append(allLinks, links...)
		pterm.DefaultBox.WithTitle(title).WithTitleTopLeft().Println(strings.Join(links, "\n"))
		pterm.Info.Printf("🔗 Found %d links from %s\n", len(links), liveDomain)
	}

	jsLinks := jsfinder.ExtractJSLinks(allLinks)
	if len(jsLinks) > 0 {
		pterm.DefaultBox.WithTitle("📜 Found JS Files").WithTitleTopLeft().Println(strings.Join(jsLinks, "\n"))
	}

	if len(jsLinks) > 0 {
		for _, js := range jsLinks {
			endpoints := jsfinder.ExtractEndpoints(js)
			title := "📌 Endpoints from: " + js
			if len(endpoints) > 0 {
				pterm.DefaultBox.WithTitle(title).WithTitleTopLeft().Println(strings.Join(endpoints, "\n"))
			} else {
				pterm.DefaultBox.WithTitle(title).WithTitleTopLeft().Println("⚠️ No endpoints found.")
			}
		}
	}

	if len(jsLinks) > 0 {
		for _, js := range jsLinks {
			secretHits := secrets.FindSecrets(js)
			title := "🔐 Secrets in: " + js
			if len(secretHits) > 0 {
				pterm.DefaultBox.WithTitle(title).WithTitleTopLeft().Println(strings.Join(secretHits, "\n"))
			} else {
				pterm.DefaultBox.WithTitle(title).WithTitleTopLeft().Println("✅ No secrets found.")
			}
		}
	}
}
