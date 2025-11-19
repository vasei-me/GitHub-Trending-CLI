package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/spf13/pflag"
)

const version = "2.0.0"

type Repo struct {
    FullName    string   `json:"full_name"`
    Description string   `json:"description"`
    Stars       int      `json:"stargazers_count"`
    Language    string   `json:"language"`
    HTMLURL     string   `json:"html_url"`
    CreatedAt   string   `json:"created_at"`
    Fork        bool     `json:"fork"`
    Archived    bool     `json:"archived"`
    Topics      []string `json:"topics"`
}

type Response struct {
    Items []Repo `json:"items"`
}

var (
    duration    = pflag.StringP("duration", "d", "week", "Time range: day, week, month, year")
    limit       = pflag.IntP("limit", "l", 10, "Number of repositories (1-100)")
    language    = pflag.String("language", "", "Filter by programming language")
    langAlias   = pflag.String("lang", "", "Alias for --language")
    jsonOutput  = pflag.Bool("json", false, "Output as JSON")
    saveFile    = pflag.String("save", "", "Save output to file")
    openBrowser = pflag.Bool("open", false, "Open first repository")
    spoken      = pflag.String("spoken", "", "Filter by spoken language")
    proxy       = pflag.String("proxy", "", "HTTP/HTTPS proxy")

    token       = pflag.String("token", os.Getenv("GITHUB_TOKEN"), "GitHub token")
    noColor     = pflag.Bool("no-color", false, "Disable colors")
    watch       = pflag.DurationP("watch", "w", 0, "Auto refresh (e.g. 5m)")
    today       = pflag.Bool("today", false, "Shortcut for --duration day")
    weekly      = pflag.Bool("weekly", false, "Shortcut for --duration week")
    monthly     = pflag.Bool("monthly", false, "Shortcut for --duration month")
    showVersion = pflag.Bool("version", false, "Show version")
)

func main() {
    pflag.Parse()

    if *noColor {
        color.NoColor = true
    }

    if *showVersion {
        color.Cyan("ghtrend v%s - GitHub Trending CLI", version)
        return
    }

    if *today {
        *duration = "day"
    }
    if *weekly {
        *duration = "week"
    }
    if *monthly {
        *duration = "month"
    }

    if *langAlias != "" && *language == "" {
        *language = *langAlias
    }
    if *limit < 1 || *limit > 100 {
        *limit = 10
    }

    if *watch > 0 {
        for {
            fetchAndShow()
            color.New(color.Faint).Printf("\nRefreshing in %s... (Ctrl+C to stop)\n", *watch)
            time.Sleep(*watch)
        } } else {
        fetchAndShow()

        color.New(color.FgCyan).Println("\nPress Enter to exit...")
        _, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
    }
}

func fetchAndShow() {
    repos, cached := fetchRepos()
    if len(repos) == 0 {
        color.Red("No data available")
        os.Exit(1)
    }

    if cached {
        color.Yellow("Using cached data")
    } else {
        color.Green("Live data • %s", time.Now().Format("02 Jan 15:04"))
        // --- شروع تغییر ---
        // خطای errcheck را برطرف کردیم
        if err := saveCache(repos); err != nil {
            color.Yellow("Warning: Failed to save cache: %v", err)
        }
    
    }

    handleOutput(repos)
}

func fetchRepos() ([]Repo, bool) {
    query := buildQuery()
    apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d",
        url.QueryEscape(query), *limit)

    client := createHTTPClient()
    req, _ := http.NewRequest("GET", apiURL, nil)
    req.Header.Set("User-Agent", "ghtrend/"+version)
    req.Header.Set("Accept", "application/vnd.github.v3+json")
    if *token != "" {
        req.Header.Set("Authorization", "token "+*token)
        color.Cyan("Using GitHub token")
    }

    resp, err := client.Do(req)
    if err != nil || resp.StatusCode != 200 {
        color.Yellow("Failed to fetch live data → using cache")
        return loadCache(), true
    }
    defer resp.Body.Close()

    var data Response
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil || len(data.Items) == 0 {
        return loadCache(), true
    }

    sort.Slice(data.Items, func(i, j int) bool {
        return dailyStars(&data.Items[i]) > dailyStars(&data.Items[j])
    })

    return data.Items, false
}

func buildQuery() string {
    daysMap := map[string]int{"day": 1, "week": 7, "month": 30, "year": 365}
    days := daysMap[strings.ToLower(*duration)]
    if days == 0 {
        days = 7
    }

    since := time.Now().AddDate(0, 0, -days).Format("2006-01-02")
    parts := []string{fmt.Sprintf("created:>%s", since)}

    if *language != "" {
        parts = append(parts, "language:"+strings.ToLower(*language))
    }
    if *spoken != "" {
        parts = append(parts, "language:"+strings.ToLower(*spoken))
    }

    parts = append(parts, "stars:>50", "fork:false")
    return strings.Join(parts, " ")
}

func dailyStars(r *Repo) int {
    if r.CreatedAt == "" {
        return r.Stars
    }
    created, err := time.Parse(time.RFC3339, r.CreatedAt)
    if err != nil {
        return r.Stars
    }
    days := int(time.Since(created).Hours()/24) + 1
    if days < 1 {
        days = 1
    }
    return r.Stars / days
}

func starsText(r *Repo) string {
    s := dailyStars(r)
    if s >= 1000 {
        return fmt.Sprintf("%.1fk/day", float64(s)/1000.0)
    }
    return fmt.Sprintf("%d/day", s)
}

func handleOutput(repos []Repo) {
    if *jsonOutput {
        b, _ := json.MarshalIndent(repos, "", "  ")
        printOrSave(string(b), "json")
        return
    }
    if *saveFile != "" {
        var sb strings.Builder
        for i, r := range repos {
            sb.WriteString(fmt.Sprintf("%d. %s – %s – %s\n   %s\n\n",
                i+1, r.FullName, starsText(&r), r.Language, r.HTMLURL))
        }
        printOrSave(sb.String(), "txt")
        return
    }
    if *openBrowser && len(repos) > 0 {
        if err := browser.OpenURL(repos[0].HTMLURL); err != nil {
            color.Red("Failed to open browser: %v", err)
        } else {
            color.Green("Opened → %s", repos[0].FullName)
        }
    }

    displayRepos(repos)
}

func displayRepos(repos []Repo) {
    fmt.Println()
    color.New(color.Bold, color.FgHiCyan).Printf("GitHub Trending • Top %d • %s\n\n",
        len(repos), time.Now().Format("02 Jan 2006"))

    for i, r := range repos {
        rank := color.YellowString("%2d", i+1)
        name := color.New(color.Bold, color.FgWhite).Sprint(r.FullName)
        stars := color.GreenString("⭐ " + starsText(&r))
        lang := languageBadge(r.Language)

        desc := r.Description
        if desc == "" {
            desc = color.New(color.Faint).Sprint("(no description)")
        }
        if len([]rune(desc)) > 90 {
            desc = string([]rune(desc)[:87]) + "..."
        }

        topics := ""
        if len(r.Topics) > 0 {
            var tags []string
            for _, t := range r.Topics {
                if len(tags) >= 4 {
                    break
                }
                tags = append(tags, color.New(color.FgBlack, color.BgWhite).Sprint(" "+t+" "))
            }
            topics = "  " + strings.Join(tags, " ")
        }

        fmt.Printf("  %s  %s  %s  %s  %s\n", rank, name, stars, lang, desc)
        fmt.Printf("     %s%s\n\n", color.BlueString(r.HTMLURL), topics)
    }
}

func languageBadge(lang string) string {
    if lang == "" {
        return color.New(color.FgHiBlack).Sprint("—")
    }

    colors := map[string]color.Attribute{
        "Go": color.FgCyan, "Python": color.FgHiBlue, "JavaScript": color.FgYellow,
        "TypeScript": color.FgHiCyan, "Rust": color.FgHiRed, "Java": color.FgHiMagenta,
        "C++": color.FgHiMagenta, "C": color.FgHiWhite, "Shell": color.FgGreen,
        "HTML": color.FgRed, "CSS": color.FgHiMagenta, "Zig": color.FgHiYellow,
    }

    attr, ok := colors[lang]
    if !ok {
        attr = color.FgMagenta
    }
    c := color.New(attr, color.Bold)
    return c.Sprint("● " + lang)
}

// کش هوشمند
func cachePath() string {
    dir, _ := os.UserCacheDir()
    return filepath.Join(dir, "ghtrend", "cache.json")
}

func saveCache(repos []Repo) error {
    type cache struct {
        Time time.Time `json:"time"`
        Data []Repo    `json:"data"`
    }
    data, err := json.Marshal(cache{time.Now(), repos})
    if err != nil {
        return err
    }
    if err := os.MkdirAll(filepath.Dir(cachePath()), 0755); err != nil {
        return err
    }
    return os.WriteFile(cachePath(), data, 0644)
}

func loadCache() []Repo {
    b, err := os.ReadFile(cachePath())
    if err != nil {
        return fallbackRepos()
    }
    var c struct {
        Time time.Time `json:"time"`
        Data []Repo    `json:"data"`
    }
    if json.Unmarshal(b, &c) != nil || time.Since(c.Time) > 15*time.Minute {
        return fallbackRepos()
    }
    return c.Data
}

func fallbackRepos() []Repo {
    color.Yellow("No cache → using built-in fallback")
    return []Repo{
        {FullName: "ollama/ollama", Description: "Run Llama 3, Mistral, Gemma locally", Stars: 85000, Language: "Go", HTMLURL: "https://github.com/ollama/ollama"},
        {FullName: "comfyanonymous/ComfyUI", Description: "Powerful Stable Diffusion GUI", Stars: 47000, Language: "Python", HTMLURL: "https://github.com/comfyanonymous/ComfyUI"},
    }
}

func createHTTPClient() *http.Client {
    tr := &http.Transport{}
    if *proxy != "" {
        if u, err := url.Parse(*proxy); err == nil {
            tr.Proxy = http.ProxyURL(u)
            color.Cyan("Proxy → %s", *proxy)
        }
    }
    return &http.Client{Timeout: 20 * time.Second, Transport: tr}
}

func printOrSave(content, ext string) {
    if *saveFile == "" {
        fmt.Print(content)
        return
    }

    fname := *saveFile
    if !strings.HasSuffix(strings.ToLower(fname), "."+ext) {
        fname += "." + ext
    }
    if err := os.WriteFile(fname, []byte(content), 0644); err != nil {
        color.Red("Failed to save file: %v", err)
    } else {
        color.Green("Saved → %s", fname)
    }
}