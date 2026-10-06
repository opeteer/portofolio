import { ProjectsPageContent } from "@/components/public/projects/projects-page-content";
import type { Metadata } from "next";

const baseUrl = process.env.NEXT_PUBLIC_SITE_URL || 'https://opeteer.dev';

export const metadata: Metadata = {
  title: "Projects",
  description: "Explore open-source systems repositories, cloud infrastructure, and tools by Gerardo M Ardianta (@opeteer).",
  keywords: ["open source", "projects", "systems engineering", "kubernetes", "go", "devops", "opeteer"],
  openGraph: {
    title: "Projects — OPETEER",
    description: "Explore open-source systems repositories, cloud infrastructure, and tools by Gerardo M Ardianta (@opeteer).",
    url: `${baseUrl}/projects`,
    type: "website",
    images: [
      {
        url: `${baseUrl}/og-image.png`,
        width: 1200,
        height: 630,
        alt: "OPETEER Projects",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Projects — OPETEER",
    description: "Explore open-source systems repositories, cloud infrastructure, and tools by Gerardo M Ardianta (@opeteer).",
    images: [`${baseUrl}/og-image.png`],
  },
  alternates: {
    canonical: `${baseUrl}/projects`,
  },
};

export default function ProjectsPage() {
  return (
    <div className="pt-24">
      <ProjectsPageContent />
    </div>
  );
}
