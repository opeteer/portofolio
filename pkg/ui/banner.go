package ui

import (
	"fmt"
	"strings"
)

const BannerASCII = `
   ____  ____  ______ _____ ______ ______ ____ 
  / __ \/ __ \/ ____// ___// ____// ____// __ \
 / / / / /_/ / __/   \__ \/ __/  / __/  / /_/ /
/ /_/ / ____/ /___  ___/ / /___ / /___ / _, _/ 
\____/_/   /_____/ /____/_____//_____//_/ |_|  
`

func PrintBanner() {
	lines := strings.Split(strings.Trim(BannerASCII, "\n"), "\n")
	for _, l := range lines {
		fmt.Println(Cyan(Bold(l)))
	}
	fmt.Println(BrightCyan("  Gerardo M Ardianta") + Gray(" — ") + White("Cloud-Native, Systems & DevOps Engineer"))
	fmt.Println(Gray("  Location: ") + Yellow("Yogyakarta, Indonesia") + Gray(" | GitHub: ") + BrightCyan("@opeteer"))
	fmt.Println(Gray("  " + strings.Repeat("─", 65)))
	fmt.Println()
}
