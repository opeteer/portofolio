package ui

import (
	"fmt"
	"strings"
)

const LogoBraille = `                                                ⢀⠄
                                          ⢠⠂   ⣠⠊
                              ⢀⡔         ⣠⠏  ⢀⡼⠁
                             ⣠⡟         ⣰⡟  ⣠⠟
         ⣿⣦⡀               ⣠⣾⡟      ⣼  ⠈⠉ ⢀⣾⠋  ⣠⠖⠁
         ⣿⡿⣿⣦⡀           ⢀⣼⣿⡟      ⣰⣿⡇   ⣴⠿⠁ ⣤⡾⠁
         ⣿⡇⠈⠻⣿⣦⡀       ⢀⣴⣿⣿⠏      ⢰⣿⣿⣧      ⠘⠋
         ⣿⡇  ⠈⠻⣿⣦⡀    ⣠⣾⣿⣿⠏      ⢠⣿⣿⣿⣿⣀⣀⣀⣀⣀⣀
         ⣿⡇    ⠈⠻⣿⡦ ⣠⣾⣿⣿⣿⢏⣠⠞⠁   ⢀⣾⣿⣿⣿⣿⣿⣿⣿⡿⠛⠁     ⢀⠄
         ⣿⡇      ⠈⢀⣴⣿⣿⣿⣿⣿⡿⠋    ⢀⣾⣿⣿⣿⣿⣿⣿⠟⠉      ⢀⣴⠋
         ⣿⡇     ⢀⣴⣿⣿⣿⣿⣿⣿⠟⠁   ⢀⣴⣿⣿⣿⣿⣿⡿⠋⠁      ⢀⣴⡿⠃
         ⣿⡇    ⣠⣾⣿⣿⣿⣿⣿⡿⢋⣤  ⣠⣶⣿⣿⣿⣿⣿⣿⠃       ⢀⣴⣿⡿⠁
         ⣿⡇  ⢠⣾⣿⣿⣿⣿⣿⣿⣿⣷⣿⠃⣠⣾⣿⣿⣿⣿⣿⣿⣿⠏      ⢀⣴⣿⣿⡟⠁
         ⣿⣧  ⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⠃⣰⣿⣿⣿⣿⣿⣿⣿⣿⡟   ⡀ ⢀⣴⣿⣿⣿⠟
         ⣿⣿   ⠈⢿⣿⣿⣿⣿⣿⣿⡏⣰⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁ ⣠⡞⢁⣴⣿⣿⣿⣿⠏
         ⣿⣿    ⠈⢿⣿⣿⣿⣿⡟⣼⣿⣿⣿⣿⣿⣿⣿⡿⠟⢁⣤⣾⡟⣴⣿⣿⣿⣿⣿⠃
         ⣿⣿     ⠈⢿⣿⣿⣿⣽⣿⣿⣿⣿⣿⣿⡿⠋⣠⣴⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁
         ⣿⣿      ⠈⢿⣿⣿⣿⣿⣿⣿⣿⡿⣫⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁
 ⢀⡀      ⣿⣿       ⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟ ⣤⡀
 ⠈⠻⣿⣿⢶⣶⣦⣤⣿⣿       ⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏ ⠈⠻⣿⣦⡀
   ⠙⢿⣦⡀⠈⠉⠉⠛       ⣾⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠋    ⠈⢻⣿⣦
     ⠻⣷⣄         ⢰⣿⠟⣛⢻⣿⣿⡿⠃              ⣠⣴⡿⠟⠁
      ⠘⢿⣧⡀       ⢹⣿⣌⣛⣰⣿⠟⠁            ⢀⣴⣾⠿⠋
        ⠙⣿⣦⡀      ⢻⣿⣿⣿⠋            ⣠⣶⡿⠛⠁
         ⠈⠻⣿⣦⡀    ⠈⢿⣿⡇          ⢀⣴⣾⠿⠋
           ⠈⠻⣷⣄    ⠘⣿         ⣠⣾⡿⠟⠁
             ⠙⢿⣷⣄   ⠈      ⣀⣴⣿⠿⠋
               ⠙⢿⣷⣄     ⢀⣠⣾⣿⠟⠁
                 ⠙⢿⣷⣄ ⣀⣴⣿⡿⠋
                   ⠙⣿⣿⣿⠟⠁
                    ⠈⠋⠁`

const LogoBlock = `                                     ▄
                        ▄        █  ▄▀
                        █       ▄█  █
       ▄               █▀       ▀  █
       ██             ██     █    ▄▀  ▄
       ███           ▄█      █   ▄█  █
       ████         ▄██     ██   ▀  █▀
       ██ ██       ▄██      ██▄     ▀
       ██  ██      ███     ▄███
       ██   ██    ███      ████▄▄▄▄
       ██    ██  ███▀      ███████▀
       ██     ▀ ██████    ███████▀     ▄
       ██      ▄█████     ██████      ▄█
       ██     ▄█████▀    █████▀       █
       ██    ▄█████▀    █████▀       █▀
       ██    ██████    █████▀       ██
       ██   ██████▄█ ▄██████       ██▀
       ██  ████████▀▄███████      ███
       ██  ████████ ███████▀     ███
       ██  ▀███████▄███████     ████
       ██   ██████ ████████  █ ████
       ██   ▀█████▄███████  █ ▄███▀
       ██    ████████████▀ ██▄████
       ██    ▀██████████▀▄███████▀
       ██     █████████ ▄████████
       ██     ▀███████▄███████▀
       ██      █████████████████
       ██      ▀███████████████ ▄
 ▀██▄▄ ██      ███████████████▀ ██
  ██▀████      ███████████████   ██
   █▄  ▀▀      ██████████████▀    ██
    █▄        ███████             ██
    ▀█        █▀ ███             ██
     ██       █ ▀▄█▀           ▄█▀
      █▄      ██▄██           ▄█▀
       █▄      ███           ██▀
       ▀█▄     ██▀          ██
        ▀█▄    ▀█         ▄█▀
         ▀█▄    █        ▄█▀
          ▀█    ▀       ██▀
           ██          ██
            ██       ▄██
             ██     ▄█▀
              ██   ██▀
               ██▄██▀
                ███▀
                ▀█`

func RenderShowtime(style, color string, jsonOutput bool) {
	if jsonOutput {
		PrintJSON(map[string]string{
			"logo_name": "opeteer_falcon_emblem",
			"type":      "ascii_art",
			"style":     style,
			"braille":   LogoBraille,
			"block":     LogoBlock,
		})
		return
	}

	art := LogoBraille
	if strings.ToLower(style) == "block" || strings.ToLower(style) == "blocks" {
		art = LogoBlock
	}

	colorFn := BrightCyan
	switch strings.ToLower(color) {
	case "white":
		colorFn = White
	case "green":
		colorFn = BrightGreen
	case "yellow":
		colorFn = BrightYellow
	case "magenta":
		colorFn = Magenta
	case "red":
		colorFn = Red
	case "cyan":
		colorFn = BrightCyan
	}

	lines := strings.Split(art, "\n")
	for _, l := range lines {
		fmt.Println("  " + colorFn(l))
	}

	fmt.Println()
	fmt.Println("  " + Dim("╭────────────────────────────────────────────────────────────╮"))
	fmt.Println("  " + Dim("│") + Bold(BrightCyan("                    GERARDO M ARDIANTA                      ")) + Dim("│"))
	fmt.Println("  " + Dim("│") + White("                 OPETEER • SYSTEMS & CLOUD                  ") + Dim("│"))
	fmt.Println("  " + Dim("│") + Italic(Yellow("    \"Crafting resilient systems from kernel to cloud\"       ")) + Dim("│"))
	fmt.Println("  " + Dim("╰────────────────────────────────────────────────────────────╯"))
	fmt.Println()
}
