# OPETEER CLI & Portfolio

> Single-binary, high-performance developer portfolio and systems explorer CLI tool for **Gerardo M Ardianta (@opeteer)**, written in **Golang**.

---

## Quick Start

### 1. Build Locally
```bash
make build
# or: go build -o opeteer .
```

### 2. Run
```bash
# Run interactive menu:
./opeteer

# Or run subcommands directly:
./opeteer bench              # Real-time systems micro-benchmark (CPU, Mem, Engine)
./opeteer live               # Real-time GitHub API latency & telemetry (non-mock)
./opeteer serve              # Launches embedded Web UI on http://localhost:8081
./opeteer showtime           # Displays high-res ASCII Art emblem
./opeteer about              # View bio & engineering focus
./opeteer flagship           # Deep showcase of flagship projects
./opeteer skills             # View technical skills & tooling
./opeteer infra              # ASCII architecture diagrams
./opeteer stats              # View repository metrics
./opeteer contact            # Contact info & social links
```

### 3. Install to PATH (Run anywhere)
```bash
make install
# or: cp opeteer $(go env GOPATH)/bin/opeteer
```
Once installed, simply run `opeteer` from any terminal directory.

---

## Available Commands

| Command | Aliases | Description |
| :--- | :--- | :--- |
| `opeteer` | `menu`, `interactive` | Launches interactive menu with ASCII banner |
| `opeteer bench` | `benchmark`, `speed` | Systems micro-benchmark (CPU crypto, Memory GC, HTTP Engine) |
| `opeteer live` | `telemetry`, `probe` | Real-time live network latency probe & GitHub telemetry |
| `opeteer serve` | `web`, `server` | Launches embedded Go HTTP web server on port **8081** |
| `opeteer showtime` | `logo`, `art` | Displays high-definition ASCII Art emblem & signature card |
| `opeteer about` | `bio`, `info` | View developer background, mission, and technical focus |
| `opeteer flagship` | `featured` | Deep showcase of flagship projects (PANDORA, hornetzDrive, etc.) |
| `opeteer projects` | `repos`, `list` | Browse all 31 public repositories with filters |
| `opeteer project <name>` | `inspect` | Detailed architecture and stack breakdown of a specific repo |
| `opeteer skills` | `stack` | Categorized technical skills & tooling matrix |
| `opeteer infra` | `arch` | ASCII architecture diagrams (zomboid-server & Kubernetes ingress) |
| `opeteer stats` | `metrics` | Real-time breakdown of languages, categories, and code volume |
| `opeteer contact` | - | Contact email (`gmayella245@gmail.com`), GitHub, and location |
| `opeteer status` | `ping` | Operational health and collaboration availability check |

---

## Flags & Filtering Examples

### Filter Repositories by Category:
```bash
./opeteer projects --category devops
./opeteer projects --category systems
./opeteer projects --category web
```

### Filter Repositories by Language:
```bash
./opeteer projects --lang go
./opeteer projects --lang python
```

### Limit Number of Results:
```bash
./opeteer projects --limit 5
```

### Machine-Readable JSON Output (for scripts / `jq`):
```bash
./opeteer live --json | jq .probe
./opeteer projects --json | jq '.[].name'
./opeteer skills --json
./opeteer stats --json
```

### Live Diagnostics & Telemetry:
```bash
./opeteer live
./opeteer live --limit 3
./opeteer live --timeout 2
```

### Systems Micro-Benchmarking:
```bash
./opeteer bench                           # Run full suite (CPU, Mem, Engine)
./opeteer bench --target cpu --duration 2s # Cryptographic SHA-256 stress
./opeteer bench --target mem              # Allocator & GC throughput
./opeteer bench --target engine           # HTTP router & JSON RPS test
./opeteer bench --json | jq .score_rating
```

### Subcommand Help:
```bash
./opeteer --help
./opeteer help projects
./opeteer help project
```
