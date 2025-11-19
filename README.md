<div align="center">

# 🚀 ghtrend: GitHub Trending CLI

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vasei-me/GitHub_Trending_CLI?label=Go&logo=go&style=flat-square)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/releases/latest)
[![GitHub all releases](https://img.shields.io/github/downloads/vasei-me/GitHub_Trending_CLI/total?style=flat-square&color=blue)](https://github.com/vasei-me/GitHub_Trending_CLI/releases)
[![GitHub stars](https://img.shields.io/github/stars/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/vasei-me/GitHub_Trending_CLI?style=flat-square&logo=github)](https://github.com/vasei-me/GitHub_Trending_CLI/network)
[![GitHub license](https://img.shields.io/github/license/vasei-me/GitHub_Trending_CLI?style=flat-square)](LICENSE)

**A fast, beautiful, and powerful command-line interface (CLI) to discover trending repositories on GitHub.**

`ghtrend` allows you to easily filter and view the newest and most popular open-source projects right in your terminal, based on various criteria like time, programming language, and star count.

[Installation & Usage](#-installation--usage) • [How to Use](#-how-to-use) • [Examples](#-examples) • [Contributing](#-contributing)

</div>

## ✨ Features

- ⚡️ **Fast & Up-to-Date**: Fetches real-time data via the official GitHub API.
- 🎨 **Colorful & Readable Output**: Displays results in a beautiful format with custom color-coding for each programming language (matching GitHub's style).
- 🔍 **Powerful Filters**:
  - Filter by **programming language** (e.g., `go`, `python`, `rust`).
  - Filter by **time range** (`day`, `week`, `month`, `year`).
  - Filter by repository's **spoken language** (e.g., `persian`, `english`).
- 🌐 **Proxy Support**: For users who don't have direct access to GitHub or need to bypass restrictions.
- 💾 **Flexible Output**:
  - Display output in the terminal.
  - Output in **JSON** format for use in scripts.
  - Save results to `txt` or `json` files.
- 🌍 **Open in Browser**: Open the top repository in your default browser with a single click.
- 📦 **Cached Output**: If the API is unreachable, it displays the last cached list to ensure you always have results.

## 🚀 Installation & Usage

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or newer) installed.

### 1. Clone the Repository

First, clone the project from GitHub and navigate into the main directory:

```bash
git clone https://github.com/vasei-me/GitHub_Trending_CLI.git
cd GitHub_Trending_CLI/ghtrend
```
