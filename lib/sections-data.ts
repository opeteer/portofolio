export interface TechSection {
  id: string
  number: string
  title: string
  subtitle: string
  description: string
  ascii: string
  specs: { label: string; value: string }[]
  commands: string[]
}

export const techSections: TechSection[] = [
  {
    id: "project-web-devops",
    number: "01",
    title: "project-web-devops",
    subtitle: "Web & DevOps Infrastructure",
    description: "Aplikasi web dengan node.js beserta dengan infrastrukturnya yang ditopang oleh docker containerization.",
    ascii: `
    ┌─────────────────────────┐
    │  DOCKER ENGINE           │
    │  ┌───────┐ ┌───────┐   │
    │  │ NODE  │ │  WEB  │   │
    │  └───┬───┘ └───┬───┘   │
    │      │         │        │
    │  ┌───┴─────────┴───┐   │
    │  │   CONTAINER OS   │   │
    │  └─────────────────┘   │
    └─────────────────────────┘`,
    specs: [
      { label: "Language", value: "JavaScript" },
      { label: "Container", value: "Docker" },
      { label: "Platform", value: "Node.js" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/project-web-devops",
      "Cloning into 'project-web-devops'...",
      "$ docker-compose up -d",
      "Starting containers... [OK]",
    ],
  },
  {
    id: "bettercap",
    number: "02",
    title: "bettercap",
    subtitle: "Network Reconnaissance",
    description: "The Swiss Army knife for 802.11, BLE, HID, CAN-bus, IPv4 and IPv6 networks reconnaissance and MITM attacks.",
    ascii: `
       [ATTACKER]
         │
    [ROUTER]───[TARGET]
         │
    [NETWORK]`,
    specs: [
      { label: "Language", value: "Go" },
      { label: "Domain", value: "Networking / Security" },
      { label: "Support", value: "IPv4/IPv6, BLE" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/bettercap",
      "Cloning into 'bettercap'...",
      "$ sudo bettercap -eval \"net.probe on\"",
      "Probing network... [OK]",
    ],
  },
  {
    id: "leadrive",
    number: "03",
    title: "LeaDrive",
    subtitle: "Driving Course Platform",
    description: "LeaDrive adalah aplikasi berbasis web untuk membantu pengguna mencari, membandingkan, dan memesan kursus mengemudi secara online.",
    ascii: `
    ┌──────────┐      ┌──────────┐
    │  USER    │─────>│ SEARCH   │
    └──────────┘      └──────────┘
         │                  │
    ┌────┴────┐        ┌────┴────┐
    │ BOOKING │<───────│ COMPARE  │
    └─────────┘        └─────────┘`,
    specs: [
      { label: "Language", value: "PHP" },
      { label: "Type", value: "Web Application" },
      { label: "Focus", value: "Course Booking" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/LeaDrive",
      "Cloning into 'LeaDrive'...",
      "$ php -S localhost:8000",
      "Development Server started",
    ],
  },
  {
    id: "image-recognition-yolo",
    number: "04",
    title: "Image-Recognition-YOLO",
    subtitle: "Computer Vision",
    description: "Image recognition dengan menggunakan YOLO pada dataset image yang ada, digunakan untuk projek UAS.",
    ascii: `
    Source Image
        │
    ┌───▼───┐
    │ YOLOvX│ ──> Bounding Boxes
    └───┬───┘
    ┌───▼────┐
    │ RENDER │ ──> Output Image
    └────────┘`,
    specs: [
      { label: "Language", value: "Jupyter Notebook" },
      { label: "Model", value: "YOLO" },
      { label: "Domain", value: "Computer Vision" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/Image-Recognition-with-YOLO-UAS",
      "Cloning repository...",
      "$ jupyter notebook",
      "Starting Jupyter server...",
    ],
  },
  {
    id: "the-one-sdu-mod",
    number: "05",
    title: "the-one-sdu-mod",
    subtitle: "Network Simulation",
    description: "Modifications and extensions for The ONE Simulator (Opportunistic Network Environment).",
    ascii: `
    Node A ──> Node B
                 │
              Node C
                 │
              Node D ──> Node E`,
    specs: [
      { label: "Language", value: "Java" },
      { label: "Domain", value: "DTN Simulation" },
      { label: "Tool", value: "The ONE Simulator" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/the-one-sdu-mod",
      "Cloning into 'the-one-sdu-mod'...",
      "$ ./compile.sh",
      "Build successful [OK]",
    ],
  },
  {
    id: "proyek-perpustakaan",
    number: "06",
    title: "Proyek_Perpustakaan",
    subtitle: "Library Management",
    description: "Proyek akhir sistem informasi perpustakaan.",
    ascii: `
        A ──┐
            ├──[BOOKS]──┐
        B ──┘           │
                        ├──[LOANS]
        C ──┐           │
            ├──[USERS]──┘
        D ──┘`,
    specs: [
      { label: "Language", value: "Java" },
      { label: "Type", value: "Information System" },
      { label: "Domain", value: "Management" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/Proyek_Perpustakaan",
      "Cloning into 'Proyek_Perpustakaan'...",
      "$ java -jar lib-sys.jar",
      "System loaded [OK]",
    ],
  },
  {
    id: "leadrive-admin",
    number: "07",
    title: "Leadrive-Admin",
    subtitle: "Admin Dashboard",
    description: "Admin panel and management interface for the LeaDrive platform.",
    ascii: `
    ┌─────────────────────────┐
    │     ADMIN DASHBOARD      │
    ├─────────────────────────┤
    │     DATA ANALYTICS       │
    ├─────────────────────────┤
    │     USER MANAGEMENT      │
    └─────────────────────────┘`,
    specs: [
      { label: "Language", value: "PHP" },
      { label: "Role", value: "Administration" },
      { label: "System", value: "LeaDrive" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/Leadrive-Admin",
      "Cloning into 'Leadrive-Admin'...",
      "$ tail -f access.log",
      "Monitoring admin access...",
    ],
  },
  {
    id: "leadrive-pemilik-kursus",
    number: "08",
    title: "LeaDrive-Pemilik",
    subtitle: "Course Owner Dashboard",
    description: "Dashboard application for driving course owners to manage their services on LeaDrive.",
    ascii: `
    ┌─────────────────────────┐
    │     COURSE MANAGEMENT    │
    ├─────────────────────────┤
    │     SCHEDULE SYSTEM      │
    ├─────────────────────────┤
    │     REVENUE TRACKING     │
    └─────────────────────────┘`,
    specs: [
      { label: "Language", value: "Blade" },
      { label: "Framework", value: "Laravel" },
      { label: "Role", value: "Partner" },
    ],
    commands: [
      "$ git clone https://github.com/gerardomayella/LeaDrive-Pemilik-Kursus",
      "Cloning into 'LeaDrive-Pemilik-Kursus'...",
      "$ php artisan serve",
      "Laravel development server started",
    ],
  }
]

export const navLinks = techSections.map((s) => ({
  id: s.id,
  number: s.number,
  title: s.title,
}))
