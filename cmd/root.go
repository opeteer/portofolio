package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

const AppVersion = "1.0.0"

func PrintHelp(topic string) {
	topic = strings.ToLower(strings.TrimSpace(topic))

	switch topic {
	case "projects", "repos", "list":
		fmt.Println(ui.Bold(ui.BrightCyan("OPETEER CLI: projects command")))
		fmt.Println()
		fmt.Println(ui.Bold("USAGE:"))
		fmt.Println("  opeteer projects [flags]")
		fmt.Println()
		fmt.Println(ui.Bold("ALIASES:"))
		fmt.Println("  repos, list")
		fmt.Println()
		fmt.Println(ui.Bold("FLAGS:"))
		fmt.Println("  -c, --category <name>   Filter by category (devops, systems, web, mobile, ai, academic)")
		fmt.Println("  -l, --lang <language>   Filter by programming language (go, python, php, java, etc.)")
		fmt.Println("  -n, --limit <number>    Limit number of results to display")
		fmt.Println("  -j, --json              Output results in raw JSON format")
		fmt.Println()
		fmt.Println(ui.Bold("EXAMPLES:"))
		fmt.Println("  $ opeteer projects")
		fmt.Println("  $ opeteer projects --category devops")
		fmt.Println("  $ opeteer projects --lang go")
		fmt.Println("  $ opeteer projects --limit 5")
		fmt.Println("  $ opeteer projects --category systems --json")
		fmt.Println()
		return

	case "project", "inspect":
		fmt.Println(ui.Bold(ui.BrightCyan("OPETEER CLI: project command")))
		fmt.Println()
		fmt.Println(ui.Bold("USAGE:"))
		fmt.Println("  opeteer project <repository-name> [flags]")
		fmt.Println()
		fmt.Println(ui.Bold("ALIASES:"))
		fmt.Println("  inspect")
		fmt.Println()
		fmt.Println(ui.Bold("FLAGS:"))
		fmt.Println("  -j, --json              Output results in raw JSON format")
		fmt.Println()
		fmt.Println(ui.Bold("EXAMPLES:"))
		fmt.Println("  $ opeteer project zomboid-server")
		fmt.Println("  $ opeteer project PANDORA")
		fmt.Println("  $ opeteer project hornetzDrive")
		fmt.Println("  $ opeteer inspect infra-for-kube --json")
		fmt.Println()
		return

	case "skills", "stack":
		fmt.Println(ui.Bold(ui.BrightCyan("OPETEER CLI: skills command")))
		fmt.Println()
		fmt.Println(ui.Bold("USAGE:"))
		fmt.Println("  opeteer skills [flags]")
		fmt.Println()
		fmt.Println(ui.Bold("FLAGS:"))
		fmt.Println("  -j, --json              Output results in raw JSON format")
		fmt.Println()
		fmt.Println(ui.Bold("EXAMPLES:"))
		fmt.Println("  $ opeteer skills")
		fmt.Println("  $ opeteer skills --json")
		fmt.Println()
		return

	case "flagship", "featured":
		fmt.Println(ui.Bold(ui.BrightCyan("OPETEER CLI: flagship command")))
		fmt.Println()
		fmt.Println(ui.Bold("USAGE:"))
		fmt.Println("  opeteer flagship [flags]")
		fmt.Println()
		fmt.Println(ui.Bold("DESCRIPTION:"))
		fmt.Println("  Displays curated showcase of top flagship projects including PANDORA,")
		fmt.Println("  hornetzDrive, infra-for-kube, zomboid-server, LeaDrive, and YOLO recognition.")
		fmt.Println()
		fmt.Println(ui.Bold("FLAGS:"))
		fmt.Println("  -j, --json              Output results in raw JSON format")
		fmt.Println()
		return

	case "showtime", "showcase", "logo", "art", "emblem":
		PrintShowtimeHelp()
		return
	}

	// Default main help
	fmt.Println(ui.Bold(ui.Cyan("OPETEER CLI")) + " - " + ui.White("Gerardo M Ardianta Portfolio & Systems Explorer"))
	fmt.Println()
	fmt.Println(ui.Bold("USAGE:"))
	fmt.Println("  opeteer [command] [flags]")
	fmt.Println("  opeteer [flags]")
	fmt.Println()
	fmt.Println(ui.Bold("CORE COMMANDS:"))
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "showtime")), "Display high-res ASCII Art emblem & logo (aliases: logo, art)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "about")), "Display bio, background, and engineering focus (aliases: bio, info)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "flagship")), "Showcase top flagship engineering projects (aliases: featured)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "projects")), "List and filter public repositories (aliases: repos, list)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "project <name>")), "Deep-dive inspection of a specific project (aliases: inspect)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "skills")), "View technical skills & tooling matrix (aliases: stack)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "infra")), "Inspect cloud & homelab architecture diagrams (aliases: arch)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "stats")), "Display GitHub repository metrics and language breakdown")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "contact")), "Show social links, email, and contact info")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "status")), "Check operational and availability status (aliases: ping)")
	fmt.Printf("  %s %s\n", ui.BrightCyan(fmt.Sprintf("%-16s", "interactive")), "Launch interactive prompt-driven navigation (aliases: ui, menu)")
	fmt.Println()
	fmt.Println(ui.Bold("GLOBAL FLAGS:"))
	fmt.Printf("  %s %s\n", fmt.Sprintf("%-16s", "-h, --help"), "Show help for opeteer or a subcommand")
	fmt.Printf("  %s %s\n", fmt.Sprintf("%-16s", "-v, --version"), "Show opeteer version information")
	fmt.Printf("  %s %s\n", fmt.Sprintf("%-16s", "-j, --json"), "Output results in raw JSON format for scripting/jq")
	fmt.Printf("  %s %s\n", fmt.Sprintf("%-16s", "    --no-color"), "Disable ANSI color formatting")
	fmt.Println()
	fmt.Println(ui.Bold("EXAMPLES:"))
	fmt.Println("  $ opeteer")
	fmt.Println("  $ opeteer flagship")
	fmt.Println("  $ opeteer projects --category devops")
	fmt.Println("  $ opeteer projects --lang go")
	fmt.Println("  $ opeteer project zomboid-server")
	fmt.Println("  $ opeteer skills --json")
	fmt.Println("  $ opeteer help projects")
	fmt.Println()
	fmt.Println("Use \"" + ui.Cyan("opeteer [command] --help") + "\" or \"" + ui.Cyan("opeteer help [command]") + "\" for detailed subcommand information.")
	fmt.Println()
}

func RunInteractive() {
	ui.PrintBanner()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(ui.Bold(ui.Cyan("Interactive Menu:")))
		fmt.Println("  " + ui.BrightCyan("1)") + " About & Engineering Mission")
		fmt.Println("  " + ui.BrightCyan("2)") + " Flagship Projects Showcase")
		fmt.Println("  " + ui.BrightCyan("3)") + " Browse All Repositories (31 projects)")
		fmt.Println("  " + ui.BrightCyan("4)") + " Technical Skills Matrix")
		fmt.Println("  " + ui.BrightCyan("5)") + " Cloud & System Architecture Diagrams")
		fmt.Println("  " + ui.BrightCyan("6)") + " GitHub Metrics & Insights")
		fmt.Println("  " + ui.BrightCyan("7)") + " Contact & Availability")
		fmt.Println("  " + ui.BrightCyan("8)") + " Inspect a Specific Project")
		fmt.Println("  " + ui.BrightCyan("9)") + " View Official ASCII Art Emblem (Showtime)")
		fmt.Println("  " + ui.BrightCyan("q)") + " Exit")
		fmt.Print(ui.Bold("\nChoose an option [1-9, q]: "))

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())
		fmt.Println()

		switch strings.ToLower(choice) {
		case "1", "about":
			RunAbout(false)
		case "2", "flagship":
			RunFlagship(false)
		case "3", "projects", "repos":
			RunProjects(ProjectFilter{}, false)
		case "4", "skills":
			RunSkills(false)
		case "5", "infra":
			RunInfra(false)
		case "6", "stats":
			RunStats(false)
		case "7", "contact":
			RunContact(false)
		case "8", "inspect":
			fmt.Print(ui.Bold("Enter repository name to inspect (e.g. zomboid-server, PANDORA): "))
			if scanner.Scan() {
				target := strings.TrimSpace(scanner.Text())
				if target != "" {
					fmt.Println()
					RunInspectProject(target, false)
				}
			}
		case "9", "showtime", "showcase", "logo", "art":
			RunShowtime(ShowtimeOptions{Style: "braille", Color: "cyan"}, false)
		case "q", "exit", "quit":
			fmt.Println(ui.BrightGreen("Thank you for visiting! Have a great day."))
			return
		default:
			fmt.Println(ui.Red("Invalid selection. Please choose 1-9 or q."))
		}

		fmt.Println(ui.Dim("Press Enter to continue..."))
		scanner.Scan()
		fmt.Println()
	}
}

func Execute(args []string) {
	var jsonOutput = false
	var noColor = false
	var showHelp = false
	var showVersion = false

	var filter ProjectFilter
	var cleanedArgs []string

	// Parse flags and arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "-j" || arg == "--json":
			jsonOutput = true
		case arg == "--no-color":
			noColor = true
			ui.NoColor = true
		case arg == "-h" || arg == "--help":
			showHelp = true
		case arg == "-v" || arg == "--version":
			showVersion = true
		case arg == "-c" || arg == "--category":
			if i+1 < len(args) {
				filter.Category = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--category="):
			filter.Category = strings.TrimPrefix(arg, "--category=")
		case arg == "-l" || arg == "--lang":
			if i+1 < len(args) {
				filter.Language = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--lang="):
			filter.Language = strings.TrimPrefix(arg, "--lang=")
		case arg == "-n" || arg == "--limit":
			if i+1 < len(args) {
				if val, err := strconv.Atoi(args[i+1]); err == nil {
					filter.Limit = val
				}
				i++
			}
		case strings.HasPrefix(arg, "--limit="):
			if val, err := strconv.Atoi(strings.TrimPrefix(arg, "--limit=")); err == nil {
				filter.Limit = val
			}
		default:
			cleanedArgs = append(cleanedArgs, arg)
		}
	}

	if noColor {
		ui.NoColor = true
	}

	if showVersion {
		fmt.Printf("opeteer version %s\n", AppVersion)
		return
	}

	if len(cleanedArgs) == 0 {
		if showHelp {
			PrintHelp("")
			return
		}
		// If running in terminal without arguments, show banner and interactive mode
		RunInteractive()
		return
	}

	cmd := strings.ToLower(cleanedArgs[0])

	if showHelp {
		PrintHelp(cmd)
		return
	}

	switch cmd {
	case "help":
		sub := ""
		if len(cleanedArgs) > 1 {
			sub = cleanedArgs[1]
		}
		PrintHelp(sub)

	case "about", "bio", "info":
		RunAbout(jsonOutput)

	case "contact":
		RunContact(jsonOutput)

	case "status", "ping":
		RunStatus(jsonOutput)

	case "flagship", "featured":
		RunFlagship(jsonOutput)

	case "showtime", "showcase", "logo", "art", "emblem":
		opts, jsonOut := ParseShowtimeFlags(args)
		if jsonOutput {
			jsonOut = true
		}
		RunShowtime(opts, jsonOut)

	case "projects", "repos", "list":
		RunProjects(filter, jsonOutput)

	case "project", "inspect":
		if len(cleanedArgs) < 2 {
			ui.Fatal("Please specify a project name to inspect. Example: opeteer project zomboid-server")
		}
		RunInspectProject(cleanedArgs[1], jsonOutput)

	case "skills", "stack":
		RunSkills(jsonOutput)

	case "infra", "architecture", "arch":
		RunInfra(jsonOutput)

	case "stats", "metrics":
		RunStats(jsonOutput)

	case "interactive", "menu", "ui":
		RunInteractive()

	case "quote":
		RunQuote()

	default:
		// Check if user directly typed a project name e.g. 'opeteer zomboid-server'
		if proj := data.FindProjectByName(cmd); proj != nil {
			RunInspectProject(proj.Name, jsonOutput)
			return
		}
		ui.Fatal("Unknown command '%s'. Run 'opeteer --help' for a list of available commands.", cmd)
	}
}
