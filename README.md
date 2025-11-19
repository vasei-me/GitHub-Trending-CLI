<div align="center">

# 🚀 ghtrend: GitHub Trending CLI

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://makeapullrequest.com)

**A blazing-fast, beautiful, and feature-rich command-line interface to discover trending repositories on GitHub right from your terminal!** ⚡

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

Method 3: Go Install
go install github.com/vasei-me/GitHub_Trending_CLI/ghtrend@latest

🎯 Quick Start

# Basic usage - get trending repos from the past week

ghtrend

# Filter by programming language

ghtrend --language go

# Get today's trending repositories

ghtrend --duration day --limit 5

📖 Usage Examples
Basic Commands
| Command | Description |
|----------------------------------------------|--------------------------------------------------------------|
| `ghtrend` | Top 10 trending repos this week (default) |
| `ghtrend --today` | Today’s trending repositories |
| `ghtrend --weekly` | This week’s trending (default) |
| `ghtrend --monthly` | This month’s trending |
| `ghtrend --duration day|week|month|year` | Custom time range |
| `ghtrend -l 25` or `ghtrend --limit 25` | Show up to 100 repositories |
| `ghtrend --language go` | Filter by language (go, python, rust, etc.) |
| `ghtrend --lang js` | Short alias for --language |
| `ghtrend --spoken english` | Filter by spoken language |
| `ghtrend --open` | Open the #1 repo in your browser |
| `ghtrend --json` | Output as JSON |
| `ghtrend --save file.txt` | Save results to a text file |
| `ghtrend --save file.json --json` | Save results as JSON file |
| `ghtrend --token ghp_xxx` | Use GitHub token (higher rate limits) |
| `ghtrend -w 5m` or `ghtrend --watch 5m` | Auto-refresh every 5 minutes (or any interval) |
| `ghtrend --proxy http://127.0.0.1:10809` | Use HTTP/HTTPS or SOCKS proxy |
| `ghtrend --no-color` | Disable colored output (good for scripts/pipes) |
| `ghtrend --version` | Show version (v2.0.0) |

Filtering Examples

# Filter by programming language

ghtrend --language python
ghtrend --lang javascript # Alias for --language
ghtrend --language go --limit 8

# Filter by time range

ghtrend --duration day # Today's trending
ghtrend --duration week # This week (default)
ghtrend --duration month # This month
ghtrend --duration year # This year

# Filter by spoken language

ghtrend --spoken persian
ghtrend --spoken english --language python

# Combined filters

ghtrend --language rust --duration month --limit 15
Output Options

# JSON output for scripting

ghtrend --language python --json

# Save to file

ghtrend --save trending.txt
ghtrend --language go --save go_repos.json --json

# Open first result in browser

ghtrend --open
ghtrend --language javascript --open --limit 1

Advanced Usage

# Using proxy

ghtrend --proxy http://127.0.0.1:10809
ghtrend --proxy socks5://127.0.0.1:10808

# Environment variable proxy

export HTTP_PROXY=http://127.0.0.1:10809
ghtrend --language python

# Complex workflow example

ghtrend --language typescript --duration day --json --save daily_ts.json

# Using proxy

ghtrend --proxy http://127.0.0.1:10809
ghtrend --proxy socks5://127.0.0.1:10808

# Environment variable proxy

export HTTP_PROXY=http://127.0.0.1:10809
ghtrend --language python

# Complex workflow example

ghtrend --language typescript --duration day --json --save daily_ts.json

🔧 Command Reference
Global Flags
| Flag | Short | Default | Description |
| ------------ | ----- | ------- | ------------------------------------------ |
| `--duration` | `-d` | `week` | Time range: `day`, `week`, `month`, `year` |
| `--limit` | `-l` | `10` | Number of repositories (1-100) |
| `--language` | | `""` | Filter by programming language |
| `--lang` | | `""` | Alias for `--language` |
| `--spoken` | | `""` | Filter by spoken language |
| `--json` | | `false` | Output as JSON format |
| `--save` | | `""` | Save output to file |
| `--open` | | `false` | Open first repo in browser |
| `--proxy` | | `""` | HTTP/HTTPS proxy URL |
| `--version` | | `false` | Show version information |

🛠️ Development
Building from Source

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

🤝 Contributing
We love your input! We want to make contributing to ghtrend as easy and transparent as possible.

How to Contribute
Fork the repo
Create a feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add some amazing feature')
Push to the branch (git push origin feature/amazing-feature)
Open a Pull Request
Reporting Issues
Please use the GitHub Issues to report bugs or suggest new features.

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

🙏 Acknowledgments
Thanks to GitHub for the amazing API
Built with Go and amazing open-source libraries
Inspired by the need for a fast, reliable trending tool

<div align="center">

⭐ Star this repo if you find it useful!
Happy coding! 🚀

</div>
```
