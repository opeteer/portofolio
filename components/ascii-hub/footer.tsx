"use client"

import { motion } from "framer-motion"
import { Github, Twitter, Linkedin, ArrowUp } from "lucide-react"

const ASCII_LOGO = `
  ____                     _       
 / ___| ___ _ __ __ _ _ __| | ___  
| |  _ / _ \\ '__/ _\` | '__| |/ _ \\ 
| |_| |  __/ | | (_| | |  | | (_) |
 \\____|\\___|_|  \\__,_|_|  |_|\\___/ 
`

const socialLinks = [
  { name: "GitHub", icon: Github, href: "https://github.com/gerardomayella" },
  { name: "Twitter", icon: Twitter, href: "https://twitter.com" },
  { name: "LinkedIn", icon: Linkedin, href: "https://linkedin.com" },
]

export function Footer() {
  const scrollToTop = () => {
    window.scrollTo({ top: 0, behavior: "smooth" })
  }

  return (
    <footer className="border-t border-border">
      <div className="mx-auto max-w-7xl px-4 py-16 lg:px-8 lg:py-24">
        <div className="grid gap-12 lg:grid-cols-3">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.5 }}
          >
            <pre
              className="font-mono text-[8px] leading-[10px] text-foreground/40 md:text-[10px] md:leading-[12px]"
              aria-label="Monochrome Hub ASCII logo"
              role="img"
            >
              {ASCII_LOGO}
            </pre>
            <p className="mt-4 max-w-xs font-mono text-xs leading-relaxed text-muted-foreground">
              Aspiring DevOps Engineer. Passionate about Linux System Administration, Networking, Cloud Infrastructure, and Automating software delivery pipelines.
            </p>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.5, delay: 0.1 }}
          >
            <span className="mb-4 block font-mono text-xs uppercase tracking-widest text-muted-foreground">
              Connect
            </span>
            <div className="flex flex-col gap-2">
              {socialLinks.map((link) => (
                <a
                  key={link.name}
                  href={link.href}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="group flex items-center gap-3 py-2 font-mono text-sm text-muted-foreground transition-all duration-200 hover:text-foreground focus-visible:ring-2 focus-visible:ring-foreground focus-visible:outline-none"
                >
                  <link.icon size={14} />
                  <span>{link.name}</span>
                  <span className="ml-auto opacity-0 transition-opacity duration-200 group-hover:opacity-100">
                    {"->"}
                  </span>
                </a>
              ))}
            </div>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.5, delay: 0.2 }}
            className="flex flex-col justify-between"
          >
            <div>
              <span className="mb-4 block font-mono text-xs uppercase tracking-widest text-muted-foreground">
                Tech Stack
              </span>
              <div className="flex flex-wrap gap-2">
                {["Linux", "Docker", "Git", "GitHub Actions", "Ansible", "Terraform", "Kubernetes", "Bash"].map(
                  (tech) => (
                    <span
                      key={tech}
                      className="border border-border px-2 py-1 font-mono text-[10px] text-muted-foreground"
                    >
                      {tech}
                    </span>
                  )
                )}
              </div>
            </div>

            <button
              onClick={scrollToTop}
              className="mt-8 flex items-center gap-2 self-start font-mono text-xs text-muted-foreground transition-all duration-200 hover:text-foreground focus-visible:ring-2 focus-visible:ring-foreground focus-visible:outline-none lg:self-end"
              aria-label="Back to top"
            >
              <ArrowUp size={12} />
              <span>BACK TO TOP</span>
            </button>
          </motion.div>
        </div>

        <div className="mt-16 flex flex-col items-center justify-between gap-4 border-t border-border pt-8 sm:flex-row">
          <span className="font-mono text-[10px] text-muted-foreground">
            {"// "} Gerardo M Ardianta &mdash; {new Date().getFullYear()}
          </span>
          <span className="font-mono text-[10px] text-muted-foreground">
            Jl. Paingan, Krodan, Maguwoharjo, Sleman, DIY
          </span>
        </div>
      </div>
    </footer>
  )
}
