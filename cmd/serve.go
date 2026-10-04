package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/server"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

func PrintServeHelp() {
	ui.PrintSectionHeader("OPETEER CLI: serve command")
	fmt.Println("USAGE:")
	fmt.Println("  opeteer serve [flags]")
	fmt.Println()
	fmt.Println("ALIASES:")
	fmt.Println("  web, server, daemon")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Launches a real embedded Go HTTP web server that serves the interactive")
	fmt.Println("  web portfolio UI, the official ASCII art emblem, and real REST API endpoints.")
	fmt.Println()
	fmt.Println("FLAGS:")
	fmt.Println("  -p, --port <number>    Port to listen on (default: 8081)")
	fmt.Println("      --host <string>    Host address to bind to (default: 0.0.0.0)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  $ opeteer serve")
	fmt.Println("  $ opeteer serve --port 8081")
	fmt.Println("  $ opeteer serve -p 3000 --host 127.0.0.1")
	fmt.Println()
}

func RunServe(args []string) {
	cfg := server.Config{
		Host: "0.0.0.0",
		Port: 8081, // Default port set to 8081 per user instruction
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "-h" || arg == "--help":
			PrintServeHelp()
			return
		case arg == "-p" || arg == "--port":
			if i+1 < len(args) {
				if p, err := strconv.Atoi(args[i+1]); err == nil && p > 0 {
					cfg.Port = p
				}
				i++
			}
		case strings.HasPrefix(arg, "--port="):
			val := strings.TrimPrefix(arg, "--port=")
			if p, err := strconv.Atoi(val); err == nil && p > 0 {
				cfg.Port = p
			}
		case arg == "--host":
			if i+1 < len(args) {
				cfg.Host = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--host="):
			cfg.Host = strings.TrimPrefix(arg, "--host=")
		}
	}

	if err := server.StartServer(cfg); err != nil && err != http.ErrServerClosed {
		ui.Fatal("Web server failed: %v", err)
	}
}
