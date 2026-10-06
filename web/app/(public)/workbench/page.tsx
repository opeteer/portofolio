import { WorkbenchPageContent } from "@/components/public/workbench/workbench-page-content";
import type { Metadata } from "next";

const baseUrl = process.env.NEXT_PUBLIC_SITE_URL || 'https://opeteer.dev';

export const metadata: Metadata = {
  title: "Workbench",
  description: "Active experiments, prototypes, and work in progress by Gerardo M Ardianta (@opeteer).",
  keywords: ["experiments", "prototypes", "work in progress", "systems", "cloud", "opeteer"],
  openGraph: {
    title: "Workbench — OPETEER",
    description: "Active experiments, prototypes, and work in progress.",
    url: `${baseUrl}/workbench`,
    type: "website",
    images: [
      {
        url: `${baseUrl}/og-image.png`,
        width: 1200,
        height: 630,
        alt: "OPETEER Workbench",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Workbench — OPETEER",
    description: "Active experiments, prototypes, and work in progress.",
    images: [`${baseUrl}/og-image.png`],
  },
  alternates: {
    canonical: `${baseUrl}/workbench`,
  },
};

export default function WorkbenchPage() {
  return (
    <div className="pt-24">
      <WorkbenchPageContent />
    </div>
  );
}
