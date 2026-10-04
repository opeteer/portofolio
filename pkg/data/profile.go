package data

type Profile struct {
	Name         string            `json:"name"`
	Username     string            `json:"username"`
	Title        string            `json:"title"`
	Bio          string            `json:"bio"`
	Location     string            `json:"location"`
	Email        string            `json:"email"`
	GitHub       string            `json:"github"`
	Portfolio    string            `json:"portfolio"`
	Status       string            `json:"status"`
	Availability string            `json:"availability"`
	Focus        []string          `json:"focus"`
	Links        map[string]string `json:"links"`
}

var MyProfile = Profile{
	Name:         "Gerardo M Ardianta",
	Username:     "opeteer",
	Title:        "Software & Cloud-Native Engineer",
	Bio:          "Designing robust, distributed, and scalable architectures. Experienced in building cloud infrastructure, container orchestration, network automation, high-performance backends in Go/Python, and cross-platform platforms.",
	Location:     "Yogyakarta, Indonesia",
	Email:        "gmayella245@gmail.com",
	GitHub:       "https://github.com/opeteer",
	Portfolio:    "https://github.com/opeteer/portofolio",
	Status:       "Operational - All systems nominal",
	Availability: "Available for Cloud, Systems & Full-Stack Collaborations",
	Focus: []string{
		"Cloud-Native & Kubernetes Infrastructure",
		"High-Performance Backend Systems (Golang & Python)",
		"Network Automation & Reverse Tunneling (Tailscale & FRP)",
		"Container Orchestration & GitOps Deployment",
		"Resilient & Ad-Hoc Disaster Management Systems",
	},
	Links: map[string]string{
		"GitHub":    "https://github.com/opeteer",
		"Email":     "mailto:gmayella245@gmail.com",
		"Portfolio": "https://github.com/opeteer/portofolio",
	},
}
