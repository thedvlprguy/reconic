# Reconic 🔍

[![Go Version](https://img.shields.io/github/go-mod/go-version/thedvlprguy/reconic?style=for-the-badge)](https://golang.org)  
[![Build Status](https://img.shields.io/github/actions/workflow/status/thedvlprguy/reconic/go.yml?branch=master&style=for-the-badge)](https://github.com/thedvlprguy/reconic)  
[![License: MIT](https://img.shields.io/github/license/thedvlprguy/reconic?style=for-the-badge)](https://github.com/thedvlprguy/reconic/blob/main/LICENSE)  

**Reconic** is a high‑performance reconnaissance CLI tool built in Go. It offers subdomain enumeration, live subdomain resolution, link crawling, JS endpoint discovery, secret detection, and more—all displayed with a sleek terminal UI powered by [`pterm`](https://github.com/pterm/pterm).

---

## 🚀 Features

- 🌐 **Subdomain Enumeration** (passive via public APIs)  
- 🟢 **Live Subdomain Checking** (DNS resolution)  
- 🔗 **URL Crawling** (extract links from live hosts)  
- 📜 **JavaScript File Detection**  
- 📌 **Endpoint Extraction from JS**  
- 🔐 **Secret Detection in JS** (API tokens, keys, etc.)  
- 🎨 **Beautiful Boxed Terminal UI**  

---

## 📥 Installation

Install via Go (recommended):

```bash
go install github.com/thedvlprguy/reconic@latest
````

Ensure the binary is in your `PATH`:
Go installs to `$GOBIN` (or `$GOPATH/bin` by default).
Add it to your `PATH`, for example:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Or build from source:

```bash
git clone https://github.com/thedvlprguy/reconic.git
cd reconic
go mod tidy
go build -o reconic main.go
```

---

## 🎬 Demo

*Will Add Later.*

---

## 🧩 Usage

```bash
reconic <domain> [-o output.txt]
```

* `-o output.txt`: Save all dicovered and live subdomains into files:

  * `output.txt` – all subdomains
  * `output_live.txt` – only live ones

**Example:**

```bash
reconic example.com -o results.txt
```

Watch Reconic output colorful, well-laid-out data in clea---

## 📂 Repository Structure

```
reconic/
├── main.go
├── go.mod
├── internal/
│   ├── subfinder/  # Subdomain enumeration logic
│   ├── resolver/   # DNS resolution for live subdomains
│   ├── crawler/    # URL crawling logic
│   ├── jsfinder/   # JS link & endpoint extraction
│   └── secrets/    # JS secret detection logic
└── README.md
```

---

## 🛠️ Requirements

* Go 1.20 or newer
* Internet access (for recon tasks)

---

## 🤝 Contributing

PRs are welcome!

1. Fork it
2. Create a new feature branch
3. Commit your changes
4. Open a PR

---

## 📝 License

Licensed under the [MIT License](LICENSE).

---

Made with ❤️ by **@thedvlprguy**
🔗 [GitHub Repo](https://github.com/thedvlprguy/reconic)
