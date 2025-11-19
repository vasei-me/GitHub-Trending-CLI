package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/spf13/pflag"
)

const version = "1.5.1"

type Repo struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Language    string `json:"language"`
	HTMLURL     string `json:"html_url"`
}

type Response struct {
	Items []Repo `json:"items"`
}

var (
	duration    = pflag.StringP("duration", "d", "week", "Time range: day, week, month, year")
	limit       = pflag.IntP("limit", "l", 10, "Number of repositories (1-100)")
	language    = pflag.String("language", "", "Filter by programming language (e.g. go, python)")
	langAlias   = pflag.String("lang", "", "Alias for --language")
	jsonOutput  = pflag.Bool("json", false, "Output as JSON")
	saveFile    = pflag.String("save", "", "Save output to file (txt or json)")
	openBrowser = pflag.Bool("open", false, "Open first repository in browser")
	spoken      = pflag.String("spoken", "", "Filter by spoken language (e.g. persian, english)")
	showVersion = pflag.Bool("version", false, "Show version and exit")
	proxy       = pflag.String("proxy", "", "HTTP/HTTPS proxy (e.g. http://127.0.0.1:10809)")
)

func main() {
	pflag.Parse()

	if *showVersion {
		color.Cyan("ghtrend v%s - GitHub Trending CLI", version)
		return
	}

	if *limit < 1 || *limit > 100 {
		*limit = 10
	}

	if *langAlias != "" && *language == "" {
		*language = *langAlias
	}

	daysMap := map[string]int{"day": 1, "week": 7, "month": 30, "year": 365}
	days, ok := daysMap[strings.ToLower(*duration)]
	if !ok {
		color.Red("Error: duration must be day, week, month or year")
		os.Exit(1)
	}

	since := time.Now().AddDate(0, 0, -days).Format("2006-01-02")
	queryParts := []string{fmt.Sprintf("created:>%s", since)}

	if *language != "" {
		queryParts = append(queryParts, fmt.Sprintf("language:%s", strings.ToLower(*language)))
	}

	if *spoken != "" {
		queryParts = append(queryParts, fmt.Sprintf("spoken-language:%s", strings.ToLower(*spoken)))
	}

	queryParts = append(queryParts, "stars:>100", "-fork:true")
	query := strings.Join(queryParts, " ")

	apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d", url.QueryEscape(query), *limit)

	color.Cyan("Fetching trending repositories (%s)...", *duration)

	client := createHTTPClient()
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", "ghtrend-cli/"+version)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if err != nil {
			color.Yellow("Warning: %v", err)
		} else {
			color.Yellow("Warning: GitHub API returned status %d", resp.StatusCode)
		}
		color.Yellow("Showing cached trending repositories...")
		color.Cyan("\nTip: Use VPN or --proxy flag to access GitHub API")
		color.Cyan("Example: ghtrend.exe --proxy http://127.0.0.1:10809\n")
		repos := getCachedRepos()
		handleOutput(repos, *duration)
		return
	}
	defer resp.Body.Close()

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil || len(data.Items) == 0 {
		color.Yellow("No live data – showing cached trending repositories:")
		repos := getCachedRepos()
		handleOutput(repos, *duration)
		return
	}

	sort.Slice(data.Items, func(i, j int) bool {
		return data.Items[i].Stars > data.Items[j].Stars
	})

	color.Green("✓ Successfully fetched %d repositories\n", len(data.Items))
	handleOutput(data.Items, *duration)
}

func createHTTPClient() *http.Client {
	transport := &http.Transport{}

	// Check for proxy flag or environment variable
	proxyURL := *proxy
	if proxyURL == "" {
		proxyURL = os.Getenv("HTTP_PROXY")
		if proxyURL == "" {
			proxyURL = os.Getenv("HTTPS_PROXY")
		}
	}

	if proxyURL != "" {
		parsedURL, err := url.Parse(proxyURL)
		if err == nil {
			transport.Proxy = http.ProxyURL(parsedURL)
			color.Cyan("Using proxy: %s", proxyURL)
		}
	}

	return &http.Client{
		Timeout:   20 * time.Second,
		Transport: transport,
	}
}

func handleOutput(repos []Repo, duration string) {
	if *jsonOutput {
		jsonData, _ := json.MarshalIndent(repos, "", "  ")
		printOrSave(string(jsonData), "json")
		return
	}

	if *saveFile != "" {
		var sb strings.Builder
		for i, r := range repos {
			sb.WriteString(fmt.Sprintf("%d. %s (%s) - %d stars\n   %s\n\n", i+1, r.FullName, r.Language, r.Stars, r.HTMLURL))
		}
		printOrSave(sb.String(), "txt")
		return
	}

	if *openBrowser && len(repos) > 0 {
		err := browser.OpenURL(repos[0].HTMLURL)
		if err == nil {
			color.Green("Opened in browser: %s", repos[0].FullName)
		}
	}

	displayRepos(repos, duration)
}

func printOrSave(content, ext string) {
	if *saveFile != "" {
		filename := *saveFile
		if !strings.HasSuffix(strings.ToLower(filename), "."+ext) {
			filename += "." + ext
		}
		err := os.WriteFile(filename, []byte(content), 0644)
		if err == nil {
			color.Green("Saved to %s", filename)
		}
	} else {
		fmt.Print(content)
	}
}

func displayRepos(repos []Repo, duration string) {
	fmt.Println()
	color.New(color.FgHiCyan, color.Bold).Printf("  Top %d trending repositories (%s)\n\n", len(repos), duration)

	for i, r := range repos {
		rank := color.YellowString("%2d", i+1)
		name := color.WhiteString(r.FullName)
		stars := color.GreenString("⭐ %d", r.Stars)
		langCol := languageColor(r.Language)

		desc := r.Description
		if desc == "" {
			desc = color.New(color.Italic).Sprint("(no description)")
		}
		if len([]rune(desc)) > 80 {
			desc = string([]rune(desc)[:77]) + "..."
		}

		fmt.Printf("  %s %s %s %s %s\n", rank, name, stars, langCol, desc)
		fmt.Printf("     %s\n\n", color.BlueString(r.HTMLURL))
	}
}

// GitHub-style language colors
func languageColor(lang string) string {
	if lang == "" {
		return color.MagentaString("—")
	}

	colors := map[string]*color.Color{
		"Go":         color.New(color.FgYellow, color.Bold),
		"Python":     color.New(color.FgBlue, color.Bold),
		"JavaScript": color.New(color.FgYellow),
		"TypeScript": color.New(color.FgCyan),
		"Java":       color.New(color.FgRed),
		"Rust":       color.New(color.FgHiRed),
		"C++":        color.New(color.FgMagenta),
		"C#":         color.New(color.FgGreen),
		"C":          color.New(color.FgHiBlue),
		"HTML":       color.New(color.FgRed),
		"CSS":        color.New(color.FgHiMagenta),
		"Shell":      color.New(color.FgHiGreen),
		"PHP":        color.New(color.FgHiBlue),
	}

	if c, ok := colors[lang]; ok {
		return c.Sprint(lang)
	}
	return color.MagentaString(lang)
}

func getCachedRepos() []Repo {
	return []Repo{
		{FullName: "ollama/ollama", Description: "Run Llama 3.2, Mistral, Gemma 2 locally", Stars: 84217, Language: "Go", HTMLURL: "https://github.com/ollama/ollama"},
		{FullName: "comfyanonymous/ComfyUI", Description: "The most powerful and modular diffusion model GUI", Stars: 46893, Language: "Python", HTMLURL: "https://github.com/comfyanonymous/ComfyUI"},
		{FullName: "valkey-io/valkey", Description: "High-performance key-value datastore (Redis fork)", Stars: 15123, Language: "C", HTMLURL: "https://github.com/valkey-io/valkey"},
		{FullName: "cline/cline", Description: "Autonomous AI agent that codes in your terminal", Stars: 12890, Language: "TypeScript", HTMLURL: "https://github.com/cline/cline"},
		{FullName: "abi/screenshot-to-code", Description: "Turn screenshots into clean code", Stars: 10123, Language: "Python", HTMLURL: "https://github.com/abi/screenshot-to-code"},
	}
}