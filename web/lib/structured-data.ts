import type { BlogPost } from './blog-data'

export function generateBlogPostStructuredData(post: BlogPost, url: string) {
  return {
    '@context': 'https://schema.org',
    '@type': 'BlogPosting',
    headline: post.title,
    description: post.excerpt,
    image: `${url}/og-images/${post.slug}.png`,
    datePublished: new Date(post.date).toISOString(),
    dateModified: new Date(post.date).toISOString(),
    author: {
      '@type': 'Person',
      name: post.author.name,
      url: 'https://github.com/opeteer',
    },
    publisher: {
      '@type': 'Person',
      name: 'Gerardo M Ardianta',
      url: 'https://opeteer.dev',
    },
    mainEntityOfPage: {
      '@type': 'WebPage',
      '@id': `${url}/blog/${post.slug}`,
    },
    articleSection: post.category,
    keywords: post.tags.join(', '),
    timeRequired: post.readTime,
  }
}

export function generateWebsiteStructuredData(url: string) {
  return {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name: 'OPETEER',
    description: "Interactive portfolio and engineering showcase by Gerardo M Ardianta (@opeteer).",
    url: url,
    author: {
      '@type': 'Person',
      name: 'Gerardo M Ardianta',
      url: 'https://github.com/opeteer',
    },
    potentialAction: {
      '@type': 'SearchAction',
      target: {
        '@type': 'EntryPoint',
        urlTemplate: `${url}/blog?search={search_term_string}`,
      },
      'query-input': 'required name=search_term_string',
    },
  }
}

export function generatePersonStructuredData() {
  return {
    '@context': 'https://schema.org',
    '@type': 'Person',
    name: 'Gerardo M Ardianta',
    url: 'https://opeteer.dev',
    image: 'https://avatars.githubusercontent.com/u/102021671?v=4',
    sameAs: [
      'https://github.com/opeteer',
      'https://linkedin.com/in/gerardo-m-ardianta',
    ],
    jobTitle: 'Systems & Cloud Engineer',
    worksFor: {
      '@type': 'Organization',
      name: 'OPETEER',
    },
  }
}

export function generateBreadcrumbStructuredData(items: Array<{ name: string; url: string }>) {
  return {
    '@context': 'https://schema.org',
    '@type': 'BreadcrumbList',
    itemListElement: items.map((item, index) => ({
      '@type': 'ListItem',
      position: index + 1,
      name: item.name,
      item: item.url,
    })),
  }
}
