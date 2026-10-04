package cmd

import (
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/ui"
)

type ShowtimeOptions struct {
	Style string
	Color string
}

func RunShowtime(opts ShowtimeOptions, jsonOutput bool) {
	style := opts.Style
	if style == "" {
		style = "braille"
	}
	color := opts.Color
	if color == "" {
		color = "cyan"
	}

	ui.PrintSectionHeader("OPETEER EMBLEM SHOWTIME")
	ui.RenderShowtime(style, color, jsonOutput)
}

func PrintShowtimeHelp() {
	ui.PrintSectionHeader("OPETEER CLI: showtime command")
	println("USAGE:")
	println("  opeteer showtime [flags]")
	println()
	println("ALIASES:")
	println("  logo, art, emblem")
	println()
	println("DESCRIPTION:")
	println("  Renders high-definition ASCII Art of the official Opeteer geometric emblem.")
	println("  Default style is high-resolution Unicode Braille dot matrix.")
	println()
	println("FLAGS:")
	println("  -s, --style <braille|block>   Select rendering style (braille or block)")
	println("      --color <name>            Color accent (cyan, white, green, yellow, magenta, red)")
	println("  -j, --json                    Output raw ASCII art in JSON format")
	println()
	println("EXAMPLES:")
	println("  $ opeteer showtime")
	println("  $ opeteer showtime --style block")
	println("  $ opeteer showtime --color green")
	println("  $ opeteer showtime --style block --color yellow")
	println()
}

func ParseShowtimeFlags(args []string) (ShowtimeOptions, bool) {
	opts := ShowtimeOptions{
		Style: "braille",
		Color: "cyan",
	}
	jsonOutput := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "-j" || arg == "--json":
			jsonOutput = true
		case arg == "-s" || arg == "--style":
			if i+1 < len(args) {
				opts.Style = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--style="):
			opts.Style = strings.TrimPrefix(arg, "--style=")
		case arg == "--color":
			if i+1 < len(args) {
				opts.Color = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--color="):
			opts.Color = strings.TrimPrefix(arg, "--color=")
		}
	}
	return opts, jsonOutput
}
