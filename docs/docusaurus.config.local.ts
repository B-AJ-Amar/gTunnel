import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'gTunnel',
  tagline: 'Fast, lightweight tunneling solution written in Go',
  favicon: 'img/favicon.ico',

  // Future flags, see https://docusaurus.io/docs/api/docusaurus-config#future
  future: {
    v4: true, // Improve compatibility with the upcoming Docusaurus v4
  },

  // Set the production url of your site here
  url: 'https://b-aj-amar.github.io',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/gTunnel/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'B-AJ-Amar', // Usually your GitHub org/user name.
  projectName: 'gTunnel', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  // Plugins for search functionality
  plugins: [
    [
      require.resolve('docusaurus-lunr-search'),
      {
        languages: ['en'], // languages to index
        indexBaseUrl: true, // whether to index the base URL
        excludeRoutes: [
          '/search', // exclude search page itself
        ],
      },
    ],
  ],

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/B-AJ-Amar/gTunnel/tree/main/docs/',
        },
        blog: {
          showReadingTime: true,
          feedOptions: {
            type: ['rss', 'atom'],
            xslt: true,
          },
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/B-AJ-Amar/gTunnel/tree/main/docs/',
          // Useful options to enforce blogging best practices
          onInlineTags: 'warn',
          onInlineAuthors: 'warn',
          onUntruncatedBlogPosts: 'warn',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: 'img/gtunnel-social-card.jpg',
    colorMode: {
      defaultMode: 'light',
      disableSwitch: false,
      respectPrefersColorScheme: true,
    },
    
    // OPTION 1: Algolia Search (like React docs)
    // Uncomment this section and comment out the plugins section above to use Algolia
    /*
    algolia: {
      // The application ID provided by Algolia
      appId: 'YOUR_APP_ID',
      
      // Public API key: it is safe to commit it
      apiKey: 'YOUR_SEARCH_API_KEY',
      
      indexName: 'gtunnel',
      
      // Optional: see doc section below
      contextualSearch: true,
      
      // Optional: Specify domains where the navigation should occur through window.location instead on history.push
      externalUrlRegex: 'external\\.com|domain\\.com',
      
      // Optional: Replace parts of the item URLs from Algolia
      replaceSearchResultPathname: {
        from: '/docs/',
        to: '/',
      },
      
      // Optional: Algolia search parameters
      searchParameters: {},
      
      // Optional: path for search page that enabled by default (`false` to disable it)
      searchPagePath: 'search',
      
      // Optional: whether the insights feature is enabled or not on Docsearch (`false` by default)
      insights: false,
    },
    */
    
    navbar: {
      title: 'gTunnel',
      logo: {
        alt: 'gTunnel Logo',
        src: 'img/logo.svg',
        srcDark: 'img/logo-dark.svg', // Optional: different logo for dark mode
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'tutorialSidebar',
          position: 'left',
          label: 'Docs',
        },
        {
          to: '/blog',
          label: 'Blog',
          position: 'left'
        },
        {
          type: 'docsVersionDropdown',
          position: 'right',
        },
        {
          type: 'search',
          position: 'right',
        },
        {
          href: 'https://github.com/B-AJ-Amar/gTunnel',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Getting Started',
              to: '/docs/intro',
            },
            {
              label: 'Installation',
              to: '/docs/installation',
            },
            {
              label: 'CLI Reference',
              to: '/docs/cli-reference',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'GitHub Discussions',
              href: 'https://github.com/B-AJ-Amar/gTunnel/discussions',
            },
            {
              label: 'Issues',
              href: 'https://github.com/B-AJ-Amar/gTunnel/issues',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Blog',
              to: '/blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/B-AJ-Amar/gTunnel',
            },
            {
              label: 'Releases',
              href: 'https://github.com/B-AJ-Amar/gTunnel/releases',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} gTunnel. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ['bash', 'go', 'yaml', 'json', 'docker'],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
