"use client"

import { useState, useMemo } from "react"
import { cn } from "@/lib/utils"
import { Github, Star, GitFork, ArrowUpRight, Search, ShieldCheck } from "lucide-react"
import { realProjects, ProjectItem } from "@/lib/projects-data"

const categories = [
  { id: "all", label: "All Repositories" },
  { id: "flagship", label: "Flagship Showcases" },
  { id: "systems", label: "Systems & Go" },
  { id: "devops", label: "Cloud & DevOps" },
  { id: "web", label: "Web Applications" },
  { id: "mobile", label: "Mobile" },
  { id: "ai", label: "AI & Vision" },
]

export function ProjectsGrid() {
  const [activeCategory, setActiveCategory] = useState("all")
  const [searchQuery, setSearchQuery] = useState("")

  const filteredProjects = useMemo(() => {
    return realProjects.filter((p) => {
      // Category filter
      let matchesCategory = true
      if (activeCategory === "flagship") {
        matchesCategory = p.status === "flagship"
      } else if (activeCategory !== "all") {
        matchesCategory = p.category === activeCategory
      }

      // Search query filter
      let matchesSearch = true
      if (searchQuery.trim() !== "") {
        const q = searchQuery.toLowerCase()
        matchesSearch =
          p.name.toLowerCase().includes(q) ||
          p.description.toLowerCase().includes(q) ||
          p.language.toLowerCase().includes(q) ||
          p.tags.some((t) => t.toLowerCase().includes(q))
      }

      return matchesCategory && matchesSearch
    })
  }, [activeCategory, searchQuery])

  return (
    <section id="projects" className="px-4 sm:px-6 py-20 sm:py-28">
      <div className="mx-auto max-w-7xl space-y-10">
        {/* Section Header with Big Typography */}
        <div className="flex flex-col md:flex-row md:items-end justify-between gap-6 pb-6 border-b border-border/40">
          <div className="space-y-3">
            <p className="font-mono text-xs uppercase tracking-[0.3em] text-primary">
              ENGINEERING ARTIFACTS // GITHUB @OPETEER
            </p>
            <h2 className="text-4xl sm:text-6xl font-black tracking-tighter uppercase text-white">
              REPOSITORIES &<br />
              <span className="text-primary font-mono text-3xl sm:text-5xl">FLAGSHIP SYSTEMS</span>
            </h2>
          </div>
          <p className="max-w-md text-sm text-muted-foreground leading-relaxed">
            Curated index of 31 public repositories ranging from low-level systems architectures and disaster-recovery protocols to Kubernetes orchestrations.
          </p>
        </div>

        {/* Filter and Search Bar */}
        <div className="flex flex-col lg:flex-row gap-4 items-stretch lg:items-center justify-between">
          <div className="flex flex-wrap gap-2">
            {categories.map((cat) => (
              <button
                key={cat.id}
                onClick={() => setActiveCategory(cat.id)}
                className={cn(
                  "px-3.5 py-1.5 rounded-lg font-mono text-xs transition-all duration-200 border",
                  activeCategory === cat.id
                    ? "bg-primary text-primary-foreground border-primary font-semibold shadow-sm shadow-primary/20"
                    : "border-border/60 bg-card/40 text-muted-foreground hover:border-primary/40 hover:text-foreground"
                )}
              >
                {cat.label}
              </button>
            ))}
          </div>

          <div className="relative min-w-[280px]">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-3.5 w-3.5 text-muted-foreground" />
            <input
              type="text"
              placeholder="Search by repo, stack, or tag..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="w-full bg-card/40 border border-border/60 rounded-lg pl-9 pr-3 py-1.5 font-mono text-xs text-foreground placeholder:text-muted-foreground focus:outline-none focus:border-primary"
            />
          </div>
        </div>

        {/* Projects Grid */}
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {filteredProjects.map((project, idx) => (
            <div
              key={project.id}
              className={cn(
                "group relative rounded-xl border bg-card/60 p-6 flex flex-col justify-between transition-all duration-300 hover:border-primary/50 hover:bg-card/80 hover:shadow-xl hover:shadow-primary/5 glass",
                project.status === "flagship" ? "border-primary/30" : "border-border/60"
              )}
            >
              <div>
                {/* Card Top Meta */}
                <div className="flex items-center justify-between mb-3 text-xs font-mono">
                  {project.status === "flagship" ? (
                    <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded bg-primary/15 text-primary border border-primary/30 font-semibold">
                      <ShieldCheck className="h-3 w-3" />
                      FLAGSHIP
                    </span>
                  ) : (
                    <span className="text-muted-foreground uppercase">{project.category}</span>
                  )}
                  <span className="text-muted-foreground">{project.language}</span>
                </div>

                {/* Project Title */}
                <h3 className="text-xl font-bold tracking-tight text-foreground mb-2 group-hover:text-primary transition-colors flex items-center justify-between">
                  <span>{project.title}</span>
                  <ArrowUpRight className="h-4 w-4 opacity-0 -translate-x-1 translate-y-1 transition-all group-hover:opacity-100 group-hover:translate-x-0 group-hover:translate-y-0 text-primary" />
                </h3>

                {/* Description */}
                <p className="text-xs text-muted-foreground leading-relaxed mb-4">
                  {project.description}
                </p>

                {/* Tech Highlights if flagship */}
                {project.highlights && (
                  <ul className="mb-4 space-y-1 text-[11px] font-mono text-muted-foreground border-l-2 border-primary/40 pl-2.5">
                    {project.highlights.map((h, i) => (
                      <li key={i}>{h}</li>
                    ))}
                  </ul>
                )}

                {/* Tags */}
                <div className="flex flex-wrap gap-1.5 mb-6">
                  {project.tags.map((tag) => (
                    <span
                      key={tag}
                      className="text-[10px] font-mono px-2 py-0.5 rounded bg-secondary/80 text-muted-foreground border border-border/40"
                    >
                      {tag}
                    </span>
                  ))}
                </div>
              </div>

              {/* Card Footer: Stars, Size, and Link */}
              <div className="pt-4 border-t border-border/40 flex items-center justify-between font-mono text-xs">
                <div className="flex items-center gap-3 text-muted-foreground">
                  <span className="flex items-center gap-1">
                    <Star className="h-3.5 w-3.5 text-primary" />
                    {project.stars}
                  </span>
                  <span className="flex items-center gap-1">
                    <GitFork className="h-3.5 w-3.5" />
                    {project.forks}
                  </span>
                  <span>{project.sizeKb} KB</span>
                </div>

                <a
                  href={project.url}
                  target="_blank"
                  rel="noreferrer"
                  className="inline-flex items-center gap-1.5 text-primary hover:text-foreground transition-colors font-semibold"
                >
                  <Github className="h-3.5 w-3.5" />
                  <span>GitHub</span>
                </a>
              </div>
            </div>
          ))}
        </div>

        {filteredProjects.length === 0 && (
          <div className="text-center py-16 font-mono text-sm text-muted-foreground">
            No repositories found matching your filter criteria.
          </div>
        )}
      </div>
    </section>
  )
}
