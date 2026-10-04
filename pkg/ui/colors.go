package ui

import (
	"fmt"
	"os"
)

var NoColor = false

func init() {
	if _, exists := os.LookupEnv("NO_COLOR"); exists {
		NoColor = true
	}
}

func colorize(code, text string) string {
	if NoColor {
		return text
	}
	return fmt.Sprintf("\033[%sm%s\033[0m", code, text)
}

func Bold(text string) string         { return colorize("1", text) }
func Dim(text string) string          { return colorize("2", text) }
func Italic(text string) string       { return colorize("3", text) }
func Underline(text string) string    { return colorize("4", text) }
func Red(text string) string          { return colorize("31", text) }
func Green(text string) string        { return colorize("32", text) }
func Yellow(text string) string       { return colorize("33", text) }
func Blue(text string) string         { return colorize("34", text) }
func Magenta(text string) string      { return colorize("35", text) }
func Cyan(text string) string         { return colorize("36", text) }
func White(text string) string        { return colorize("37", text) }
func Gray(text string) string         { return colorize("90", text) }
func BrightCyan(text string) string   { return colorize("96", text) }
func BrightGreen(text string) string  { return colorize("92", text) }
func BrightYellow(text string) string { return colorize("93", text) }
