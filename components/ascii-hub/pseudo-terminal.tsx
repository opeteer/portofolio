"use client"

import { useState, useRef, useEffect, type KeyboardEvent } from "react"
import { motion } from "framer-motion"

const COMMANDS: Record<string, string[]> = {
  help: [
    "Available commands:",
    "  help       - Show this message",
    "  projects   - List all projects",
    "  inspect    - Inspect the current system",
    "  about      - About Gerardo M Ardianta",
    "  skills     - Show tech stack and skills",
    "  clear      - Clear terminal",
  ],
  projects: [
    "01  project-web-devops",
    "02  bettercap",
    "03  LeaDrive",
    "04  Image-Recognition-YOLO",
    "05  the-one-sdu-mod",
    "06  Proyek_Perpustakaan",
    "07  Leadrive-Admin",
    "08  LeaDrive-Pemilik",
  ],
  inspect: [
    "Profile: Gerardo M Ardianta",
    "Username: gerardomayella",
    "Role: Aspiring DevOps Engineer",
    "Status: OPERATIONAL",
  ],
  about: [
    "Gerardo M Ardianta",
    "",
    "Passionate about Linux System Administration, Networking,",
    "Cloud Infrastructure, and Automating software delivery pipelines.",
    "",
    "Currently focusing on mastering DevOps practices such as Infrastructure as Code,",
    "Configuration Management & Server Automation, Container Orchestration,",
    "GitOps Continuous Delivery, and Advanced Secrets & Security Management.",
  ],
  skills: [
    "OS & Networking: Linux, SSH, Cisco/GNS3, DTN",
    "Simulation: The ONE Simulator",
    "Containers: Docker, Docker Compose",
    "CI/CD: GitHub Actions, Git",
    "Languages: Java, PHP, Bash, JS, HTML, Go, Python",
  ],
}

interface TerminalLine {
  type: "input" | "output"
  content: string
}

export function PseudoTerminal() {
  const [lines, setLines] = useState<TerminalLine[]>([
    { type: "output", content: 'Welcome to gerardomayella-hub Terminal v1.0.0' },
    { type: "output", content: 'Type "help" for available commands.' },
    { type: "output", content: "" },
  ])
  const [input, setInput] = useState("")
  const scrollRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight
    }
  }, [lines])

  const processCommand = (cmd: string) => {
    const trimmed = cmd.trim().toLowerCase()
    const baseLines: TerminalLine[] = [
      ...lines,
      { type: "input", content: `$ ${cmd}` },
    ]

    if (trimmed === "clear") {
      setLines([])
      setInput("")
      return
    }

    const newLines: TerminalLine[] = [...baseLines]
    const response = COMMANDS[trimmed]
    if (response) {
      response.forEach((line) => {
        newLines.push({ type: "output", content: line })
      })
    } else if (trimmed === "") {
    } else {
      newLines.push({ type: "output", content: `command not found: ${trimmed}` })
      newLines.push({ type: "output", content: 'Type "help" for available commands.' })
    }

    newLines.push({ type: "output", content: "" })
    setLines(newLines)
    setInput("")
  }

  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      processCommand(input)
    }
  }

  return (
    <motion.section
      initial={{ opacity: 0, y: 20 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true }}
      transition={{ duration: 0.5 }}
      className="mx-auto max-w-7xl px-4 py-16 lg:px-8 lg:py-24"
    >
      <div className="mb-8 flex flex-col gap-4">
        <div className="flex items-center gap-4">
          <span className="font-mono text-sm text-muted-foreground">{">"}</span>
          <div className="h-[1px] w-12 bg-border" />
          <span className="font-mono text-xs uppercase tracking-widest text-muted-foreground">
            Interactive
          </span>
        </div>
        <h2 className="font-pixel-line text-3xl font-bold tracking-tight text-foreground md:text-5xl">
          Terminal
        </h2>
        <p className="max-w-prose font-mono text-sm leading-relaxed text-muted-foreground">
          Explore the system. Type commands to interact with the Hub.
        </p>
      </div>

      <div
        className="border border-border"
        onClick={() => inputRef.current?.focus()}
        role="application"
        aria-label="Interactive pseudo-terminal"
      >
        <div className="flex items-center gap-2 border-b border-border px-4 py-2.5">
          <div className="h-2.5 w-2.5 bg-foreground" />
          <div className="h-2.5 w-2.5 bg-muted-foreground/50" />
          <div className="h-2.5 w-2.5 bg-muted-foreground/30" />
          <span className="ml-2 font-mono text-xs text-muted-foreground">
            gerardomayella-hub ~ interactive
          </span>
        </div>

        <div
          ref={scrollRef}
          className="h-80 overflow-y-auto bg-secondary/20 p-4"
        >
          {lines.map((line, i) => (
            <div
              key={i}
              className={`font-mono text-xs leading-relaxed ${
                line.type === "input"
                  ? "text-foreground"
                  : "text-muted-foreground"
              }`}
            >
              {line.content || "\u00A0"}
            </div>
          ))}

          <div className="relative flex items-center font-mono text-xs text-foreground">
            <span className="mr-1">{"$"}</span>
            <span>{input}</span>
            <span className="animate-blink">{"█"}</span>
            <input
              ref={inputRef}
              type="text"
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={handleKeyDown}
              className="absolute inset-0 h-full w-full cursor-default border-none bg-transparent opacity-0 outline-none"
              aria-label="Terminal input"
              autoComplete="off"
              autoCorrect="off"
              spellCheck={false}
            />
          </div>
        </div>
      </div>
    </motion.section>
  )
}
