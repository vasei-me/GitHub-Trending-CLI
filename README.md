````markdown
<div align="center">

# 🚀 ghtrend: GitHub Trending CLI

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vasei-me/GitHub_Trending_CLI?label=Go&logo=go&style=flat-square)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/releases/latest)
[![GitHub all releases](https://img.shields.io/github/downloads/vasei-me/GitHub_Trending_CLI/total?style=flat-square&color=blue)](https://github.com/vasei-me/GitHub_Trending_CLI/releases)
[![GitHub stars](https://img.shields.io/github/stars/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/network)
[![GitHub license](https://img.shields.io/github/license/vasei-me/GitHub_Trending_CLI?style=flat-square)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://makeapullrequest.com)

**A blazing-fast, beautiful, and feature-rich command-line interface to discover trending repositories on GitHub right from your terminal!** ⚡

[Installation](#-installation) • [Quick Start](#-quick-start) • [Usage Examples](#-usage-examples) • [Contributing](#-contributing)

![ghtrend demo](https://via.placeholder.com/800x400/2D3748/FFFFFF?text=ghtrend+in+action+-+Colorful+terminal+output+with+GitHub+trending+repos)

</div>

## ✨ Features

| Feature                    | Description                                                     |
| -------------------------- | --------------------------------------------------------------- |
| ⚡ **Real-time Data**      | Fetches live data directly from GitHub's official API           |
| 🎨 **Beautiful Output**    | Color-coded terminal output with GitHub-style language colors   |
| 🔧 **Advanced Filtering**  | Filter by language, time range, spoken language, and star count |
| 🌐 **Proxy Support**       | Built-in proxy support for restricted networks                  |
| 💾 **Multiple Formats**    | Terminal display, JSON output, file export (TXT/JSON)           |
| 🌍 **Browser Integration** | Open top repositories directly in your browser                  |
| 🛡️ **Fallback System**     | Smart caching when API is unavailable                           |
| 📱 **Cross-Platform**      | Works on Windows, macOS, and Linux                              |

## 🚀 Installation

### Method 1: Using Pre-built Binaries (Recommended)

1. **Download the latest release from the [Releases Page](https://github.com/vasei-me/GitHub_Trending_CLI/releases)**
2. **Extract the binary and add it to your PATH**

### Method 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/vasei-me/GitHub_Trending_CLI.git
cd GitHub_Trending_CLI/ghtrend

# Build the binary
go build -o ghtrend main.go

# Move to system PATH (Linux/macOS)
sudo mv ghtrend /usr/local/bin/

# Or add to PATH manually (Windows)
# Add the directory containing ghtrend.exe to your System PATH
```
````

### Method 3: Go Install

```bash
go install github.com/vasei-me/GitHub_Trending_CLI/ghtrend@latest
```

## 🎯 Quick Start

```bash
# Basic usage - get trending repos from the past week
ghtrend

# Filter by programming language
ghtrend --language go

# Get today's trending repositories
ghtrend --duration day --limit 5
```

## 📖 Usage Examples

### Basic Commands

| Command                    | Description                           |
| -------------------------- | ------------------------------------- |
| `ghtrend`                  | Get top 10 trending repos (past week) |
| `ghtrend --limit 20`       | Get top 20 repositories               |
| `ghtrend --duration month` | Trending repos from past month        |
| `ghtrend --version`        | Show version information              |

### Filtering Examples

```bash
# Filter by programming language
ghtrend --language python
ghtrend --lang javascript          # Alias for --language
ghtrend --language go --limit 8

# Filter by time range
ghtrend --duration day            # Today's trending
ghtrend --duration week           # This week (default)
ghtrend --duration month          # This month
ghtrend --duration year           # This year

# Filter by spoken language
ghtrend --spoken persian
ghtrend --spoken english --language python

# Combined filters
ghtrend --language rust --duration month --limit 15
```

### Output Options

```bash
# JSON output for scripting
ghtrend --language python --json

# Save to file
ghtrend --save trending.txt
ghtrend --language go --save go_repos.json --json

# Open first result in browser
ghtrend --open
ghtrend --language javascript --open --limit 1
```

### Advanced Usage

```bash
# Using proxy
ghtrend --proxy http://127.0.0.1:10809
ghtrend --proxy socks5://127.0.0.1:10808

# Environment variable proxy
export HTTP_PROXY=http://127.0.0.1:10809
ghtrend --language python

# Complex workflow example
ghtrend --language typescript --duration day --json --save daily_ts.json
```

## 🔧 Command Reference

### Global Flags

| Flag         | Short | Default | Description                                |
| ------------ | ----- | ------- | ------------------------------------------ |
| `--duration` | `-d`  | `week`  | Time range: `day`, `week`, `month`, `year` |
| `--limit`    | `-l`  | `10`    | Number of repositories (1-100)             |
| `--language` |       | `""`    | Filter by programming language             |
| `--lang`     |       | `""`    | Alias for `--language`                     |
| `--spoken`   |       | `""`    | Filter by spoken language                  |
| `--json`     |       | `false` | Output as JSON format                      |
| `--save`     |       | `""`    | Save output to file                        |
| `--open`     |       | `false` | Open first repo in browser                 |
| `--proxy`    |       | `""`    | HTTP/HTTPS proxy URL                       |
| `--version`  |       | `false` | Show version information                   |

## 🛠️ Development

### Building from Source

```bash
# Clone and setup
git clone https://github.com/vasei-me/GitHub_Trending_CLI.git
cd GitHub_Trending_CLI/ghtrend

# Install dependencies
go mod download

# Build for current platform
go build -o ghtrend main.go

# Build for multiple platforms
GOOS=windows GOARCH=amd64 go build -o ghtrend-windows-amd64.exe main.go
GOOS=linux GOARCH=amd64 go build -o ghtrend-linux-amd64 main.go
GOOS=darwin GOARCH=amd64 go build -o ghtrend-darwin-amd64 main.go
```

### Running Tests

```bash
# Run the application in development
go run main.go --language go --limit 3

# Test JSON output
go run main.go --json --limit 2
```

## 🤝 Contributing

We love your input! We want to make contributing to ghtrend as easy and transparent as possible.

### How to Contribute

1. **Fork the repo**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add some amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Reporting Issues

Please use the [GitHub Issues](https://github.com/vasei-me/GitHub_Trending_CLI/issues) to report bugs or suggest new features.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to **GitHub** for the amazing API
- Built with **Go** and amazing open-source libraries
- Inspired by the need for a fast, reliable trending tool

---

<div align="center">

### **⭐ Star this repo if you find it useful!**

**Happy coding!** 🚀

</div>
```
