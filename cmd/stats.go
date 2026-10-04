package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

type StatsReport struct {
	TotalRepos      int            `json:"total_repos"`
	FlagshipCount   int            `json:"flagship_count"`
	TotalSizeKB     int            `json:"total_size_kb"`
	Languages       map[string]int `json:"languages"`
	Categories      map[string]int `json:"categories"`
	PrimaryLanguage string         `json:"primary_language"`
}

func RunStats(jsonOutput bool) {
	langs := make(map[string]int)
	cats := make(map[string]int)
	totalSize := 0
	flagshipCount := 0

	for _, p := range data.AllProjects {
		totalSize += p.SizeKB
		if p.Flagship {
			flagshipCount++
		}
		lang := p.Language
		if lang == "" {
			lang = "Other"
		}
		langs[lang]++
		cats[p.Category]++
	}

	report := StatsReport{
		TotalRepos:      len(data.AllProjects),
		FlagshipCount:   flagshipCount,
		TotalSizeKB:     totalSize,
		Languages:       langs,
		Categories:      cats,
		PrimaryLanguage: "Go / Python / Cloud",
	}

	if jsonOutput {
		ui.PrintJSON(report)
		return
	}

	ui.PrintSectionHeader("GITHUB REPOSITORY METRICS & INSIGHTS")

	fmt.Printf("  %-22s: %s\n", ui.Bold("Total Repositories"), ui.BrightCyan(fmt.Sprintf("%d repos", report.TotalRepos)))
	fmt.Printf("  %-22s: %s\n", ui.Bold("Flagship Showcases"), ui.Yellow(fmt.Sprintf("%d projects", report.FlagshipCount)))
	fmt.Printf("  %-22s: %s\n", ui.Bold("Total Codebase Size"), ui.White(fmt.Sprintf("%d KB (~%.1f MB)", report.TotalSizeKB, float64(report.TotalSizeKB)/1024.0)))
	fmt.Println()

	fmt.Println(ui.Bold(ui.Cyan("Domain / Category Breakdown:")))
	for cat, count := range cats {
		bar := strings.Repeat("■", count*2)
		fmt.Printf("  %-14s [%2d] %s\n", ui.White(cat), count, ui.BrightCyan(bar))
	}
	fmt.Println()

	fmt.Println(ui.Bold(ui.Cyan("Languages Distribution:")))
	type langPair struct {
		name  string
		count int
	}
	var pairs []langPair
	for l, c := range langs {
		pairs = append(pairs, langPair{l, c})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	for _, lp := range pairs {
		bar := strings.Repeat("■", lp.count*2)
		fmt.Printf("  %-18s [%2d] %s\n", ui.White(lp.name), lp.count, ui.Yellow(bar))
	}
	fmt.Println()
}
