package cmd

import (
	"fmt"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

type ProjectFilter struct {
	Category string
	Language string
	Limit    int
}

func RunFlagship(jsonOutput bool) {
	flagship := data.GetFlagshipProjects()
	if jsonOutput {
		ui.PrintJSON(flagship)
		return
	}

	ui.PrintSectionHeader("FLAGSHIP ENGINEERING SHOWCASE")

	for i, p := range flagship {
		fmt.Printf("%d. %s  %s\n", i+1, ui.Bold(ui.BrightCyan(p.Name)), ui.Dim("["+p.Language+"]"))
		fmt.Printf("   %s\n", ui.White(p.Description))
		if len(p.TechStack) > 0 {
			fmt.Printf("   %s %s\n", ui.Bold(ui.Gray("Stack:")), ui.Yellow(strings.Join(p.TechStack, " • ")))
		}
		if len(p.Highlights) > 0 {
			for _, h := range p.Highlights {
				fmt.Printf("   %s %s\n", ui.Cyan("•"), ui.Gray(h))
			}
		}
		fmt.Printf("   %s %s\n\n", ui.Bold(ui.Gray("Repo:")), ui.Dim(p.URL))
	}

	fmt.Println(ui.Dim("Tip: Inspect any project deeply with: ") + ui.Cyan("opeteer project <name>"))
	fmt.Println()
}

func RunProjects(filter ProjectFilter, jsonOutput bool) {
	var filtered []data.Project

	catFilter := strings.ToLower(strings.TrimSpace(filter.Category))
	langFilter := strings.ToLower(strings.TrimSpace(filter.Language))

	for _, p := range data.AllProjects {
		if catFilter != "" && !strings.Contains(strings.ToLower(p.Category), catFilter) {
			continue
		}
		if langFilter != "" && !strings.Contains(strings.ToLower(p.Language), langFilter) {
			continue
		}
		filtered = append(filtered, p)
	}

	if filter.Limit > 0 && len(filtered) > filter.Limit {
		filtered = filtered[:filter.Limit]
	}

	if jsonOutput {
		ui.PrintJSON(filtered)
		return
	}

	ui.PrintSectionHeader(fmt.Sprintf("PUBLIC REPOSITORIES (%d Found)", len(filtered)))

	fmt.Printf("%-3s %-26s %-16s %-12s %-45s\n",
		ui.Bold("#"), ui.Bold("NAME"), ui.Bold("LANGUAGE"), ui.Bold("CATEGORY"), ui.Bold("DESCRIPTION"))
	fmt.Println(ui.Gray(strings.Repeat("─", 105)))

	for i, p := range filtered {
		var paddedName string
		if p.Flagship {
			paddedName = ui.BrightCyan(fmt.Sprintf("%-26s [flagship]", p.Name))
		} else {
			paddedName = ui.White(fmt.Sprintf("%-37s", p.Name))
		}

		desc := p.Description
		if desc == "" {
			desc = "(No description provided)"
		}
		if len(desc) > 42 {
			desc = desc[:39] + "..."
		}

		fmt.Printf("%-3d %s %-16s %-12s %-45s\n",
			i+1,
			paddedName,
			ui.Yellow(fmt.Sprintf("%-16s", p.Language)),
			ui.Dim(fmt.Sprintf("%-12s", p.Category)),
			ui.Gray(desc),
		)
	}
	fmt.Println()
	fmt.Println(ui.Dim("[flagship] Indicates flagship project. Run 'opeteer project <name>' to view full details."))
	fmt.Println()
}

func RunInspectProject(name string, jsonOutput bool) {
	p := data.FindProjectByName(name)
	if p == nil {
		// Try case-insensitive search
		for _, proj := range data.AllProjects {
			if strings.EqualFold(proj.Name, name) {
				p = &proj
				break
			}
		}
	}

	if p == nil {
		ui.Fatal("Project '%s' not found. Run 'opeteer projects' to list all available repositories.", name)
		return
	}

	if jsonOutput {
		ui.PrintJSON(p)
		return
	}

	ui.PrintSectionHeader("PROJECT INSPECTION: " + p.Name)

	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Repository")), ui.BrightCyan(p.Name))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Category")), ui.Yellow(p.Category))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Language")), ui.Green(p.Language))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "GitHub URL")), ui.White(p.URL))
	fmt.Printf("  %s: %d KB\n", ui.Bold(fmt.Sprintf("%-14s", "Size")), p.SizeKB)
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Last Updated")), p.Updated)
	if p.Flagship {
		fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Status")), ui.BrightGreen("Flagship Engineering Showcase"))
	}
	fmt.Println()

	fmt.Println(ui.Bold(ui.Cyan("Description:")))
	if p.Description != "" {
		fmt.Println("  " + p.Description)
	} else {
		fmt.Println("  (No description provided)")
	}
	fmt.Println()

	if len(p.TechStack) > 0 {
		fmt.Println(ui.Bold(ui.Cyan("Technologies & Stack:")))
		fmt.Println("  " + ui.Yellow(strings.Join(p.TechStack, "  •  ")))
		fmt.Println()
	}

	if len(p.Highlights) > 0 {
		fmt.Println(ui.Bold(ui.Cyan("Key Engineering Highlights:")))
		for _, h := range p.Highlights {
			fmt.Printf("  %s %s\n", ui.BrightCyan("•"), h)
		}
		fmt.Println()
	}
}
