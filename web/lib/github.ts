import { realProjects, ProjectItem } from "./projects-data"

export interface GitHubUserProfile {
  login: string
  name: string
  avatar_url: string
  bio: string
  location: string
  public_repos: number
  followers: number
  following: number
  html_url: string
}

export interface LiveTelemetryData {
  latencyMs: number
  protocol: string
  rateRemaining: number
  rateLimit: number
  connected: boolean
}

export async function fetchGitHubProfile(): Promise<GitHubUserProfile> {
  try {
    const res = await fetch("https://api.github.com/users/opeteer", {
      next: { revalidate: 3600 },
      headers: {
        "User-Agent": "opeteer-portfolio-web/1.0",
        Accept: "application/vnd.github.v3+json",
      },
    })
    if (res.ok) {
      return await res.json()
    }
  } catch {
    // Graceful fallback
  }

  return {
    login: "opeteer",
    name: "Gerardo M Ardianta",
    avatar_url: "https://avatars.githubusercontent.com/u/127300451?v=4",
    bio: "Software & Cloud-Native Engineer | High-performance Go backends, Kubernetes orchestration, and resilient systems.",
    location: "Yogyakarta, Indonesia",
    public_repos: 31,
    followers: 4,
    following: 12,
    html_url: "https://github.com/opeteer",
  }
}

export async function fetchGitHubProjects(): Promise<ProjectItem[]> {
  try {
    const res = await fetch("https://api.github.com/users/opeteer/repos?sort=pushed&per_page=100", {
      next: { revalidate: 1800 },
      headers: {
        "User-Agent": "opeteer-portfolio-web/1.0",
        Accept: "application/vnd.github.v3+json",
      },
    })

    if (res.ok) {
      const liveRepos = await res.json()
      if (Array.isArray(liveRepos) && liveRepos.length > 0) {
        // Map live repos onto realProjects with live stars/forks updated
        return realProjects.map((p) => {
          const match = liveRepos.find((r: { name: string }) => r.name.toLowerCase() === p.name.toLowerCase())
          if (match) {
            return {
              ...p,
              stars: match.stargazers_count ?? p.stars,
              forks: match.forks_count ?? p.forks,
              sizeKb: match.size ?? p.sizeKb,
              url: match.html_url ?? p.url,
            }
          }
          return p
        })
      }
    }
  } catch {
    // Fallback to local dataset
  }

  return realProjects
}
