package data

type SkillCategory struct {
	Domain string   `json:"domain"`
	Skills []string `json:"skills"`
}

var TechnicalSkills = []SkillCategory{
	{
		Domain: "Languages",
		Skills: []string{"Go (Golang)", "Python", "TypeScript", "JavaScript", "Java", "PHP", "Bash", "SQL"},
	},
	{
		Domain: "Cloud, DevOps & Infrastructure",
		Skills: []string{"Kubernetes", "Docker", "Docker Compose", "Vagrant", "GitOps", "Tailscale Mesh", "Fast Reverse Proxy (FRP)", "Linux (Ubuntu/Debian)", "Nginx", "Systemd Services"},
	},
	{
		Domain: "Frameworks & Platforms",
		Skills: []string{"React Native", "Next.js", "Laravel", "Node.js (Express)", "Tailwind CSS", "Bootstrap"},
	},
	{
		Domain: "Engineering Practices & Tools",
		Skills: []string{"Microservices Architecture", "RESTful API Design", "Git & GitHub Actions", "YOLO Computer Vision", "PostgreSQL", "MySQL", "Makefile Pipelines", "VS Code", "Postman"},
	},
}

const ZomboidArchitecture = `
[Client (Internet)]
       │
       ▼ (UDP 16261 / 16262)
┌──────────────────────────────────────────────┐
│ Azure Cloud Server (Public IP Gateway)       │
│  - Fast Reverse Proxy Server (FRPS)          │
│  - Public Port Forwarding & DDoS Shield      │
└──────────────────────┬───────────────────────┘
                       │ Secure FRP Reverse Tunnel
                       ▼
┌──────────────────────────────────────────────┐
│ Local Host: Windows Virtualization           │
│  ┌────────────────────────────────────────┐  │
│  │ Vagrant VM: Ubuntu Focal64             │  │
│  │  ┌───────────────────────────────────┐ │  │
│  │  │ Docker Network:                   │ │  │
│  │  │  ├── FRP Client (frpc)            │ │  │
│  │  │  └── Project Zomboid (LinuxGSM)   │ │  │
│  │  └───────────────────────────────────┘ │  │
│  └────────────────────────────────────────┘  │
└──────────────────────────────────────────────┘
* Benefits: Bypasses ISP CGNAT/Firewalls, 0 Port-Forwarding, Zero Hosting Hardware Cost
`

const KubeArchitecture = `
[Public Traffic / Internet]
       │
       ▼
┌──────────────────────────────────────────────┐
│ NGINX Ingress Controller                     │
│  ├── /api        ──► [Backend Service] (ClusterIP)
│  │                     └── [Backend Pods x3]
│  ├── /           ──► [Frontend Service] (ClusterIP)
│  │                     └── [Web App Pods x2]
│  └── /metrics    ──► [Prometheus Monitoring Service]
└──────────────────────────────────────────────┘
`
