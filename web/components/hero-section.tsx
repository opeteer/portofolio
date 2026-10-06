"use client"

import Link from "next/link"
import { useEffect, useState } from "react"
import { Github, ArrowUpRight, Terminal } from "lucide-react"

const roles = [
  "DISTRIBUTED SYSTEMS",
  "KUBERNETES & GITOPS",
  "HIGH-THROUGHPUT GOLANG",
  "CONTAINER NETWORKING",
  "DISASTER RECOVERY ARCHITECTURE",
]

const officialEmblem = `                                                ⢀⠄
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

export function HeroSection() {
  const [currentRole, setCurrentRole] = useState(0)
  const [displayText, setDisplayText] = useState("")
  const [isDeleting, setIsDeleting] = useState(false)

  useEffect(() => {
    const targetText = roles[currentRole]
    const timeout = setTimeout(
      () => {
        if (!isDeleting) {
          if (displayText.length < targetText.length) {
            setDisplayText(targetText.slice(0, displayText.length + 1))
          } else {
            setTimeout(() => setIsDeleting(true), 2500)
          }
        } else {
          if (displayText.length > 0) {
            setDisplayText(displayText.slice(0, -1))
          } else {
            setIsDeleting(false)
            setCurrentRole((prev) => (prev + 1) % roles.length)
          }
        }
      },
      isDeleting ? 40 : 80,
    )
    return () => clearTimeout(timeout)
  }, [displayText, isDeleting, currentRole])

  return (
    <section className="relative px-4 sm:px-6 pt-24 sm:pt-32 pb-16 sm:pb-24 border-b border-border/40">
      <div className="mx-auto max-w-7xl">
        <div className="grid gap-12 lg:grid-cols-12 lg:gap-8 items-center min-h-[75vh]">
          {/* Left column - Big Typography */}
          <div className="lg:col-span-7 space-y-8 animate-fade-in-up">
            <div className="inline-flex items-center gap-2 rounded-full border border-primary/30 bg-primary/10 px-3.5 py-1.5 font-mono text-xs text-primary backdrop-blur-sm">
              <span className="h-2 w-2 rounded-full bg-primary animate-ping" />
              <span>OPETEER CORE // SYSTEMS OPERATIONAL</span>
            </div>

            <div className="space-y-4">
              <h1 className="text-5xl sm:text-7xl lg:text-8xl font-black tracking-tighter uppercase text-white leading-[0.9] text-balance">
                GERARDO M
                <br />
                <span className="text-primary font-mono tracking-tight text-4xl sm:text-6xl lg:text-7xl block mt-2">
                  ARDIANTA
                </span>
              </h1>

              <div className="pt-2 font-mono text-sm sm:text-base text-primary flex items-center gap-2">
                <span className="text-muted-foreground">$ focus --stream :</span>
                <span className="bg-primary/20 text-primary px-2 py-0.5 rounded border border-primary/30 font-semibold tracking-wide typing-cursor">
                  {displayText}
                </span>
              </div>
            </div>

            <p className="max-w-xl text-base sm:text-lg leading-relaxed text-muted-foreground">
              Crafting resilient distributed systems, low-latency backends in <strong className="text-foreground">Golang</strong> & <strong className="text-foreground">Python</strong>, Kubernetes orchestration pipelines, and ad-hoc disaster recovery architectures from kernel to cloud.
            </p>

            {/* Quick Metrics Strip */}
            <div className="grid grid-cols-3 gap-4 pt-2 border-t border-border/40 max-w-lg font-mono">
              <div>
                <div className="text-2xl sm:text-3xl font-black text-primary">31</div>
                <div className="text-xs text-muted-foreground uppercase">Public Repos</div>
              </div>
              <div>
                <div className="text-2xl sm:text-3xl font-black text-foreground">07</div>
                <div className="text-xs text-muted-foreground uppercase">Flagships</div>
              </div>
              <div>
                <div className="text-2xl sm:text-3xl font-black text-primary">97K+</div>
                <div className="text-xs text-muted-foreground uppercase">Engine RPS</div>
              </div>
            </div>

            <div className="flex flex-wrap gap-4 pt-2">
              <a
                href="#projects"
                className="group relative inline-flex items-center justify-center gap-3 overflow-hidden rounded-lg border border-primary bg-primary px-6 py-3.5 font-mono text-sm font-semibold text-primary-foreground shadow-lg shadow-primary/20 transition-all duration-300 hover:scale-[1.02] active:scale-[0.98]"
              >
                <span>EXPLORE REPOSITORIES</span>
                <span className="transition-transform duration-300 group-hover:translate-x-1">→</span>
              </a>

              <a
                href="https://github.com/opeteer"
                target="_blank"
                rel="noreferrer"
                className="inline-flex items-center justify-center gap-2 rounded-lg border border-border bg-card/60 px-6 py-3.5 font-mono text-sm text-foreground transition-all duration-300 hover:border-primary/50 hover:bg-card hover:text-primary active:scale-[0.98]"
              >
                <Github className="h-4 w-4" />
                <span>GITHUB @OPETEER</span>
                <ArrowUpRight className="h-3.5 w-3.5 opacity-60" />
              </a>
            </div>
          </div>

          {/* Right column - Official ASCII Art Emblem Terminal */}
          <div className="lg:col-span-5 relative animate-scale-in">
            <div className="relative rounded-xl border border-primary/30 bg-[#070b07]/90 p-5 sm:p-6 shadow-2xl shadow-primary/10 glass backdrop-blur-md">
              {/* Terminal header */}
              <div className="flex items-center justify-between pb-3 mb-4 border-b border-border/50 font-mono text-xs text-muted-foreground">
                <div className="flex items-center gap-2">
                  <div className="h-3 w-3 rounded-full bg-red-500/80" />
                  <div className="h-3 w-3 rounded-full bg-yellow-500/80" />
                  <div className="h-3 w-3 rounded-full bg-primary/80" />
                  <span className="ml-2 text-foreground font-semibold flex items-center gap-1.5">
                    <Terminal className="h-3.5 w-3.5 text-primary" />
                    opeteer://emblem.ascii
                  </span>
                </div>
                <span className="text-[10px] text-primary bg-primary/10 px-2 py-0.5 rounded border border-primary/20">
                  BRAILLE HD
                </span>
              </div>

              {/* ASCII Emblem */}
              <pre className="overflow-x-auto font-mono text-[9px] sm:text-[10.5px] leading-[1.05] tracking-tight text-primary select-none whitespace-pre py-1">
                {officialEmblem}
              </pre>

              <div className="mt-4 pt-3 border-t border-border/40 flex items-center justify-between font-mono text-xs">
                <span className="text-muted-foreground">ID: GEOMETRIC SWIFT EMBLEM</span>
                <span className="text-primary font-semibold">STATUS: NOMINAL</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
