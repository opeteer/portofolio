package cmd

import (
	"fmt"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

func RunAbout(jsonOutput bool) {
	if jsonOutput {
		ui.PrintJSON(data.MyProfile)
		return
	}

	ui.PrintSectionHeader("ABOUT THE DEVELOPER")
	p := data.MyProfile

	fmt.Printf("%s: %s (%s)\n", ui.Bold(ui.White("Name")), ui.BrightCyan(p.Name), ui.Cyan("@"+p.Username))
	fmt.Printf("%s: %s\n", ui.Bold(ui.White("Role")), ui.Yellow(p.Title))
	fmt.Printf("%s: %s\n", ui.Bold(ui.White("Location")), ui.White(p.Location))
	fmt.Printf("%s: %s\n", ui.Bold(ui.White("Email")), ui.Green(p.Email))
	fmt.Printf("%s: %s\n", ui.Bold(ui.White("Status")), ui.BrightGreen(p.Status))
	fmt.Println()

	fmt.Println(ui.Bold(ui.Cyan("Bio & Mission:")))
	fmt.Println("  " + p.Bio)
	fmt.Println()

	fmt.Println(ui.Bold(ui.Cyan("Engineering Focus:")))
	for _, f := range p.Focus {
		fmt.Printf("  %s %s\n", ui.BrightCyan("•"), f)
	}
	fmt.Println()
}

func RunContact(jsonOutput bool) {
	if jsonOutput {
		ui.PrintJSON(data.MyProfile.Links)
		return
	}

	ui.PrintSectionHeader("CONTACT & SOCIAL LINKS")
	p := data.MyProfile

	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Email")), ui.BrightCyan(p.Email))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "GitHub")), ui.White(p.GitHub))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Portfolio")), ui.White(p.Portfolio))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Location")), ui.Yellow(p.Location))
	fmt.Printf("  %s: %s\n", ui.Bold(fmt.Sprintf("%-14s", "Availability")), ui.BrightGreen(p.Availability))
	fmt.Println()
}

func RunStatus(jsonOutput bool) {
	statusData := map[string]string{
		"status":       "HEALTHY_OPERATIONAL",
		"availability": data.MyProfile.Availability,
		"location":     data.MyProfile.Location,
		"contact":      data.MyProfile.Email,
	}

	if jsonOutput {
		ui.PrintJSON(statusData)
		return
	}

	fmt.Println(ui.BrightGreen("PONG: System operational."))
	fmt.Println("  " + ui.White("Status:       ") + ui.Green(statusData["status"]))
	fmt.Println("  " + ui.White("Availability: ") + ui.Cyan(statusData["availability"]))
	fmt.Println("  " + ui.White("Location:     ") + ui.Yellow(statusData["location"]))
	fmt.Println("  " + ui.White("Contact:      ") + ui.BrightCyan(statusData["contact"]))
	fmt.Println()
}

func RunQuote() {
	quotes := []string{
		"\"Simplicity is prerequisite for reliability.\" — Edsger W. Dijkstra",
		"\"Clean code always looks like it was written by someone who cares.\" — Robert C. Martin",
		"\"Premature optimization is the root of all evil.\" — Donald Knuth",
	}
	fmt.Println(ui.Italic(ui.Cyan(quotes[0])))
	fmt.Println("  " + strings.Repeat("─", 50))
}
