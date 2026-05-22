"use client"

import { useState, useEffect } from "react"
import { Navigation } from "@/components/ascii-hub/navigation"
import { HeroSection } from "@/components/ascii-hub/hero-section"
import { DomainSection } from "@/components/ascii-hub/domain-section"
import { TechTicker } from "@/components/ascii-hub/tech-ticker"
import { PseudoTerminal } from "@/components/ascii-hub/pseudo-terminal"
import { Footer } from "@/components/ascii-hub/footer"
import { techSections } from "@/lib/sections-data"

export default function Home() {
  const [isMounted, setIsMounted] = useState(false)

  useEffect(() => {
    setIsMounted(true)
  }, [])

  return (
    <div className="min-h-screen bg-background text-foreground">
      <Navigation />

      <main>
        <HeroSection />

        <TechTicker />

        {isMounted && techSections.map((section, index) => (
          <DomainSection
            key={section.id}
            section={section}
            index={index}
          />
        ))}

        <PseudoTerminal />
      </main>

      <Footer />
    </div>
  )
}
