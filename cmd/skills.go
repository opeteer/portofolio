package cmd

import (
	"fmt"
	"strings"

	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

func RunSkills(jsonOutput bool) {
	if jsonOutput {
		ui.PrintJSON(data.TechnicalSkills)
		return
	}

	ui.PrintSectionHeader("TECHNICAL SKILLS & TOOLING MATRIX")

	for _, cat := range data.TechnicalSkills {
		fmt.Println(ui.Bold(ui.BrightCyan("┌─ " + cat.Domain)))
		for _, s := range cat.Skills {
			fmt.Printf("│  %s %s\n", ui.Yellow("•"), ui.White(s))
		}
		fmt.Println(ui.Dim("└" + strings.Repeat("─", 50)))
		fmt.Println()
	}
}

func RunInfra(jsonOutput bool) {
	infraData := map[string]string{
		"zomboid_dedicated_server": data.ZomboidArchitecture,
		"kubernetes_ingress_stack": data.KubeArchitecture,
	}

	if jsonOutput {
		ui.PrintJSON(infraData)
		return
	}

	ui.PrintSectionHeader("SYSTEM ARCHITECTURE & INFRASTRUCTURE TOPOLOGY")

	fmt.Println(ui.Bold(ui.BrightCyan("1. Hybrid Game Server Architecture (zomboid-server)")))
	fmt.Println(ui.Dim("   Vagrant VM + Containerized LinuxGSM + FRP Reverse Tunnel to Azure Cloud"))
	fmt.Println(ui.Cyan(data.ZomboidArchitecture))
	fmt.Println()

	fmt.Println(ui.Bold(ui.BrightCyan("2. Kubernetes Ingress & Microservices Topology (infra-for-kube)")))
	fmt.Println(ui.Dim("   Declarative Ingress Controller, Pod Services & Monitoring Routing"))
	fmt.Println(ui.Cyan(data.KubeArchitecture))
	fmt.Println()
}
