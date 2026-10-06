import { BlogHero } from "@/components/public/blog/blog-hero";
import { BlogList } from "@/components/public/blog/blog-list";
import { BlogSidebar } from "@/components/public/blog/blog-sidebar";
import type { Metadata } from "next";

const baseUrl = process.env.NEXT_PUBLIC_SITE_URL || 'https://opeteer.dev';

export const metadata: Metadata = {
  title: "Blog",
  description: "Technical articles, experiments, and insights on systems programming, cloud infrastructure, and Go by Gerardo M Ardianta (@opeteer).",
  openGraph: {
    title: "Blog — OPETEER",
    description: "Technical articles, experiments, and insights from the digital laboratory.",
    url: `${baseUrl}/blog`,
    type: "website",
    images: [
      {
        url: `${baseUrl}/og-image.png`,
        width: 1200,
        height: 630,
        alt: "OPETEER Blog",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Blog — OPETEER",
    description: "Technical articles, experiments, and insights from the digital laboratory.",
    images: [`${baseUrl}/og-image.png`],
  },
  alternates: {
    canonical: `${baseUrl}/blog`,
  },
};

export default function BlogPage() {
  return (
    <div>
      <BlogHero />
      <section className="px-4 sm:px-6 py-16 sm:py-20 border-t border-border/30">
        <div className="mx-auto max-w-7xl">
          <div className="grid gap-12 lg:grid-cols-[1fr_320px]">
            <BlogList />
            <BlogSidebar />
          </div>
        </div>
      </section>
    </div>
  );
}
