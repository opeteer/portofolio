export interface ProjectItem {
  id: number
  title: string
  name: string
  description: string
  tags: string[]
  category: "systems" | "devops" | "web" | "mobile" | "ai" | "academic" | "competition"
  status: "active" | "flagship" | "shipped" | "research"
  language: string
  stars: number
  forks: number
  sizeKb: number
  url: string
  homepage?: string
  featured: boolean
  highlight?: boolean
  techStack?: string[]
  highlights?: string[]
}

export const realProjects: ProjectItem[] = [
  {
    id: 1,
    title: "PANDORA",
    name: "PANDORA",
    description: "Mobile application prototype engineered for FIT COMPETITION 2026 TRACK II: Kegagalan Sistem Komunikasi Seluler. Focuses on disaster-resilient localized networking and communication failure recovery.",
    tags: ["React Native", "Java", "Mobile", "Disaster-Recovery", "Competition"],
    category: "competition",
    status: "flagship",
    language: "React Native / Java",
    stars: 3,
    forks: 1,
    sizeKb: 832,
    url: "https://github.com/opeteer/PANDORA",
    featured: true,
    highlight: true,
    techStack: ["React Native", "Java", "Local Mesh", "Offline-First"],
    highlights: [
      "Offline-first disaster mesh communication protocol",
      "Automated fallback to localized peer beacons",
      "FIT Competition 2026 Track II candidate entry"
    ]
  },
  {
    id: 2,
    title: "hornetzDrive",
    name: "hornetzDrive",
    description: "High-performance distributed cloud storage and virtual drive engine written in Golang, delivering high IOPS, fast concurrent block uploads, and low memory footprint.",
    tags: ["Go", "Distributed-Storage", "Cloud", "High-Throughput", "Systems"],
    category: "systems",
    status: "flagship",
    language: "Go",
    stars: 5,
    forks: 2,
    sizeKb: 1395,
    url: "https://github.com/opeteer/hornetzDrive",
    featured: true,
    highlight: true,
    techStack: ["Golang", "Concurrent I/O", "Chunking Protocol", "REST API"],
    highlights: [
      "Microsecond concurrent chunk serialization",
      "Zero external runtime dependencies",
      "High-throughput object buffer management"
    ]
  },
  {
    id: 3,
    title: "ztatic-go-framework",
    name: "ztatic-go-framework",
    description: "Modular high-performance HTTP web framework engineered from scratch in Go, providing ultra-fast routing, lightweight middlewares, and zero external dependency footprint.",
    tags: ["Go", "Framework", "HTTP-Engine", "Microservices", "Systems"],
    category: "systems",
    status: "flagship",
    language: "Go",
    stars: 4,
    forks: 1,
    sizeKb: 21,
    url: "https://github.com/opeteer/ztatic-go-framework",
    featured: true,
    highlight: true,
    techStack: ["Golang", "HTTP Router", "Trie Node Matching", "Zero-Alloc"],
    highlights: [
      "Custom trie-based radix router engine",
      "Zero-allocation context multiplexing",
      "Sub-millisecond middleware execution pipeline"
    ]
  },
  {
    id: 4,
    title: "infra-for-kube",
    name: "infra-for-kube",
    description: "Production-grade Kubernetes infrastructure blueprints featuring declarative manifests, NGINX Ingress routing, GitOps deployment automation, and Prometheus observability.",
    tags: ["Kubernetes", "DevOps", "Ingress", "Prometheus", "GitOps", "YAML"],
    category: "devops",
    status: "flagship",
    language: "YAML",
    stars: 6,
    forks: 2,
    sizeKb: 6,
    url: "https://github.com/opeteer/infra-for-kube",
    featured: true,
    highlight: true,
    techStack: ["Kubernetes", "NGINX Ingress", "Prometheus", "Helm"],
    highlights: [
      "Declarative microservices ingress controllers",
      "Integrated Prometheus metrics scraping",
      "Self-healing multi-pod deployment manifests"
    ]
  },
  {
    id: 5,
    title: "zomboid-server",
    name: "zomboid-server",
    description: "Dedicated production gaming infrastructure running containerized LinuxGSM inside a Vagrant Ubuntu VM, tunneled through an Azure FRP reverse proxy for complete IP masking without router port-forwarding.",
    tags: ["DevOps", "Vagrant", "Docker", "FRP", "Tailscale", "Azure", "Linux"],
    category: "devops",
    status: "flagship",
    language: "Shell / Lua",
    stars: 8,
    forks: 3,
    sizeKb: 28,
    url: "https://github.com/opeteer/zomboid-server",
    featured: true,
    highlight: true,
    techStack: ["Vagrant", "Docker Compose", "FRP Reverse Tunnel", "Azure VM"],
    highlights: [
      "Complete public IP masking via cloud reverse proxy",
      "Zero router port-forwarding required",
      "Automated LinuxGSM backup and update daemon"
    ]
  },
  {
    id: 6,
    title: "LeaDrive",
    name: "LeaDrive",
    description: "Secure cross-platform file management and document storage application built with React Native and modern cloud storage integrations.",
    tags: ["React Native", "TypeScript", "Cloud-Storage", "Mobile"],
    category: "mobile",
    status: "flagship",
    language: "TypeScript",
    stars: 2,
    forks: 0,
    sizeKb: 651,
    url: "https://github.com/opeteer/LeaDrive",
    featured: true,
    techStack: ["React Native", "TypeScript", "Encrypted Cache", "REST"],
    highlights: [
      "Cross-platform mobile document sync",
      "Local encrypted file cache",
      "Responsive file explorer UI"
    ]
  },
  {
    id: 7,
    title: "Image-Recognition-with-YOLO-UAS",
    name: "Image-Recognition-with-YOLO-UAS",
    description: "Computer vision and deep learning project utilizing the YOLO (You Only Look Once) architecture for real-time object detection and recognition datasets.",
    tags: ["Python", "YOLO", "Computer-Vision", "Deep-Learning", "AI"],
    category: "ai",
    status: "flagship",
    language: "Python / Jupyter",
    stars: 3,
    forks: 1,
    sizeKb: 4335,
    url: "https://github.com/opeteer/Image-Recognition-with-YOLO-UAS",
    featured: true,
    techStack: ["Python", "YOLOv8", "OpenCV", "PyTorch"],
    highlights: [
      "Custom dataset annotation and model training",
      "Real-time video frame object inference",
      "High-confidence bounding box evaluation"
    ]
  },
  {
    id: 8,
    title: "portofolio",
    name: "portofolio",
    description: "Developer portfolio and single-binary systems explorer CLI & Web application engineered in Go and Next.js.",
    tags: ["Go", "Next.js", "CLI", "Web", "Telemetry"],
    category: "systems",
    status: "active",
    language: "Go / TypeScript",
    stars: 5,
    forks: 1,
    sizeKb: 1024,
    url: "https://github.com/opeteer/portofolio",
    featured: true,
  },
  {
    id: 9,
    title: "opeteer",
    name: "opeteer",
    description: "Configuration files and automation workflows for Gerardo M Ardianta's GitHub profile and developer environment.",
    tags: ["Markdown", "Config", "GitHub"],
    category: "systems",
    status: "active",
    language: "Markdown",
    stars: 2,
    forks: 0,
    sizeKb: 12,
    url: "https://github.com/opeteer/opeteer",
    featured: false,
  },
  {
    id: 10,
    title: "KOPMA-TC",
    name: "KOPMA-TC",
    description: "Full-stack cooperative enterprise management and point-of-sale platform built with PHP and MySQL.",
    tags: ["PHP", "MySQL", "Web", "Enterprise"],
    category: "web",
    status: "shipped",
    language: "PHP",
    stars: 1,
    forks: 0,
    sizeKb: 12480,
    url: "https://github.com/opeteer/KOPMA-TC",
    featured: false,
  },
  {
    id: 11,
    title: "sistem-informasi-desa",
    name: "sistem-informasi-desa",
    description: "Village governance administration and public citizen records information portal.",
    tags: ["Laravel", "PHP", "Bootstrap", "Web"],
    category: "web",
    status: "shipped",
    language: "PHP",
    stars: 1,
    forks: 0,
    sizeKb: 8940,
    url: "https://github.com/opeteer/sistem-informasi-desa",
    featured: false,
  },
  {
    id: 12,
    title: "monitoring-iot-sensors",
    name: "monitoring-iot-sensors",
    description: "Real-time telemetry and IoT sensor data collector pipeline with time-series database integration.",
    tags: ["Python", "IoT", "Telemetry", "MQTT"],
    category: "systems",
    status: "shipped",
    language: "Python",
    stars: 2,
    forks: 0,
    sizeKb: 2150,
    url: "https://github.com/opeteer/monitoring-iot-sensors",
    featured: false,
  }
]
