# OPETEER CLI & Portfolio

> Single-binary, high-performance developer portfolio and systems explorer CLI tool for **Gerardo M Ardianta (@opeteer)**, written in **Golang**.

---

## ⚡ Quick Start

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
./opeteer showtime
./opeteer about
./opeteer flagship
./opeteer skills
./opeteer infra
./opeteer stats
./opeteer contact
```

### 3. Install to PATH (Run anywhere)
```bash
make install
# or: cp opeteer $(go env GOPATH)/bin/opeteer
```
Once installed, simply run `opeteer` from any terminal directory.

---

## 📖 Available Commands

| Command | Aliases | Description |
| :--- | :--- | :--- |
| `opeteer` | `menu`, `interactive` | Launches interactive menu with ASCII banner |
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

## 🛠️ Flags & Filtering Examples

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
./opeteer projects --json | jq '.[].name'
./opeteer skills --json
./opeteer stats --json
```

### Subcommand Help:
```bash
./opeteer --help
./opeteer help projects
./opeteer help project
```
