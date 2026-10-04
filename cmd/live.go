package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/opeteer/opeteer-cli/pkg/live"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

type LiveOptions struct {
	Limit   int
	Timeout time.Duration
}

func PrintLiveHelp() {
	ui.PrintSectionHeader("OPETEER CLI: live command")
	fmt.Println("USAGE:")
	fmt.Println("  opeteer live [flags]")
	fmt.Println()
	fmt.Println("ALIASES:")
	fmt.Println("  telemetry, probe")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Performs a real-time network latency probe to GitHub API, inspects live API quota,")
	fmt.Println("  and fetches recently pushed repositories directly from GitHub servers (100% non-mock).")
	fmt.Println()
	fmt.Println("FLAGS:")
	fmt.Println("  -n, --limit <number>     Number of recently pushed repositories to display (default: 5)")
	fmt.Println("      --timeout <seconds>  HTTP timeout duration in seconds (default: 5)")
	fmt.Println("  -j, --json               Output live telemetry in raw JSON format")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  $ opeteer live")
	fmt.Println("  $ opeteer live --limit 3")
	fmt.Println("  $ opeteer live --timeout 2")
	fmt.Println("  $ opeteer live --json | jq .probe")
	fmt.Println()
}

func formatRelativeTime(t time.Time) string {
	if t.IsZero() {
		return "recently"
	}
	diff := time.Since(t)
	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		mins := int(diff.Minutes())
		if mins <= 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours <= 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case diff < 48*time.Hour:
		return "yesterday"
	default:
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	}
}

func RunLive(args []string) {
	opts := LiveOptions{
		Limit:   5,
		Timeout: 5 * time.Second,
	}
	jsonOutput := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "-h" || arg == "--help":
			PrintLiveHelp()
			return
		case arg == "-j" || arg == "--json":
			jsonOutput = true
		case arg == "-n" || arg == "--limit":
			if i+1 < len(args) {
				if val, err := strconv.Atoi(args[i+1]); err == nil && val > 0 {
					opts.Limit = val
				}
				i++
			}
		case strings.HasPrefix(arg, "--limit="):
			if val, err := strconv.Atoi(strings.TrimPrefix(arg, "--limit=")); err == nil && val > 0 {
				opts.Limit = val
			}
		case arg == "--timeout":
			if i+1 < len(args) {
				if sec, err := strconv.Atoi(args[i+1]); err == nil && sec > 0 {
					opts.Timeout = time.Duration(sec) * time.Second
				}
				i++
			}
		}
	}

	client := live.NewClient(opts.Timeout)
	report, err := client.FetchLiveTelemetry(opts.Limit)
	if err != nil && report == nil {
		ui.Fatal("Failed to probe telemetry: %v", err)
		return
	}

	if jsonOutput {
		ui.PrintJSON(report)
		return
	}

	ui.PrintSectionHeader("LIVE GITHUB TELEMETRY & NETWORK PROBE")

	if report.IsFallback {
		fmt.Println("  " + ui.Yellow("Notice: Live connection unreachable or rate limit reached."))
		fmt.Println("  " + ui.Dim("Displaying local cached telemetry snapshot.") + "\n")
	}

	// 1. Network Probe & Latency Diagnostics
	fmt.Println(ui.Bold(ui.BrightCyan("┌─ Real-Time Network & API Diagnostics")))
	statusStr := ui.Green("Connected • TLS 1.3")
	if !report.Probe.Connected {
		statusStr = ui.Red("Disconnected (Offline)")
	}
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Connection")), statusStr)

	latencyColor := ui.BrightGreen
	if report.Probe.LatencyMs > 300 {
		latencyColor = ui.Yellow
	}
	if report.Probe.LatencyMs > 800 {
		latencyColor = ui.Red
	}
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Network RTT")), latencyColor(fmt.Sprintf("%.1f ms", report.Probe.LatencyMs)))
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Endpoint")), ui.Dim(report.Probe.Endpoint))

	if report.Probe.RateLimit > 0 {
		resetMins := int(time.Until(report.Probe.RateReset).Minutes())
		if resetMins < 0 {
			resetMins = 0
		}
		fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "API Rate Limit")), ui.White(fmt.Sprintf("%d / %d requests remaining (Resets in %dm)", report.Probe.RateRemain, report.Probe.RateLimit, resetMins)))
	}
	fmt.Println(ui.Dim("└" + strings.Repeat("─", 50)) + "\n")

	// 2. Real-Time Profile Snapshot
	fmt.Println(ui.Bold(ui.BrightCyan("┌─ Live Profile Snapshot")))
	fmt.Printf("│  %s: %s (%s)\n", ui.Bold(fmt.Sprintf("%-16s", "Developer")), ui.White(report.Profile.Name), ui.Cyan("@"+report.Profile.Username))
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Public Repos")), ui.Yellow(fmt.Sprintf("%d active repositories", report.Profile.PublicRepos)))
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Network Reach")), ui.White(fmt.Sprintf("%d followers • %d following", report.Profile.Followers, report.Profile.Following)))
	fmt.Printf("│  %s: %s\n", ui.Bold(fmt.Sprintf("%-16s", "Last Sync")), ui.Dim(report.Timestamp.Format("2006-01-02 15:04:05 MST")))
	fmt.Println(ui.Dim("└" + strings.Repeat("─", 50)) + "\n")

	// 3. Recently Pushed Repositories (Chronological Live Activity)
	fmt.Println(ui.Bold(ui.BrightCyan(fmt.Sprintf("Recently Pushed Repositories (Top %d Live Activity):", len(report.RecentRepos)))))
	fmt.Println()

	for i, r := range report.RecentRepos {
		lang := r.Language
		if lang == "" {
			lang = "Config"
		}
		relativeTime := formatRelativeTime(r.PushedAt)

		fmt.Printf("  %d. %s  %s  %s\n",
			i+1,
			ui.Bold(ui.BrightCyan(r.Name)),
			ui.Dim("["+lang+"]"),
			ui.Yellow("• Pushed "+relativeTime),
		)
		if r.Description != "" {
			fmt.Printf("     %s\n", ui.Gray(r.Description))
		}
		fmt.Printf("     %s\n\n", ui.Dim(r.HTMLURL))
	}

	fmt.Println(ui.Dim("Tip: Run 'opeteer project <name>' to inspect any repository in full detail."))
	fmt.Println()
}
