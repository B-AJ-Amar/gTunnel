import type {ReactNode} from 'react';
import {useState} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import ScrollNavbar from '@site/src/components/ScrollNavbar';
import { 
  Globe, 
  Shield, 
  Code, 
  Zap, 
  Container, 
  Wrench,
  Users,
  Webhook,
  Target,
  FlaskConical
} from 'lucide-react';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  const [copied, setCopied] = useState(false);
  
  const installCommand = 'curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash';
  
  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(installCommand);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
  };

  return (
    <header className={clsx('hero', styles.heroBanner)}>
      <div className="container">
        <div className={styles.heroContent}>
          <div className={styles.heroText}>
            <div className={styles.logoContainer}>
              <img 
                src="/img/logo.png" 
                alt="gTunnel Logo" 
                className={styles.heroLogo}
              />
            </div>
            <Heading as="h1" className={styles.heroTitle}>
              Fast & Secure
              <br />
              <span className={styles.highlight}>Dev Tunnel</span>
            </Heading>
            <p className={styles.heroSubtitle}>
              Expose your local development servers to the internet instantly. 
              Built with Go for maximum performance and security.
            </p>
            <div className={styles.buttons}>
              <Link
                className="button button--primary button--lg"
                to="/quick-start">
                Get Started
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" className={styles.buttonIcon}>
                  <path d="M5 12h14M12 5l7 7-7 7" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
                </svg>
              </Link>
              <Link
                className="button button--secondary button--lg button--outline"
                to="/docs/intro">
                Documentation
              </Link>
            </div>
            <div className={styles.quickDemo}>
              <div className={styles.codeBlock}>
                <div className={styles.codeHeader}>
                  <span className={styles.codeLabel}>Quick Install</span>
                  <button 
                    className={styles.copyButton}
                    onClick={copyToClipboard}
                    title={copied ? 'Copied!' : 'Copy to clipboard'}
                  >
                    {copied ? (
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M20 6L9 17L4 12" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
                      </svg>
                    ) : (
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <rect x="9" y="9" width="13" height="13" rx="2" ry="2" stroke="currentColor" strokeWidth="2" fill="none"/>
                        <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1" stroke="currentColor" strokeWidth="2" fill="none"/>
                      </svg>
                    )}
                  </button>
                </div>
                <code>
                  <span className={styles.command}>curl</span>{' '}
                  <span className={styles.flag}>-sSL</span>{' '}
                  <span className={styles.url}>https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh</span>{' '}
                  <span className={styles.operator}>|</span>{' '}
                  <span className={styles.command}>bash</span>
                </code>
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}

function FeatureCard({icon: Icon, title, description}: {icon: any, title: string, description: string}) {
  return (
    <div className={styles.featureCard}>
      <div className={styles.featureIconContainer}>
        <Icon className={styles.featureIcon} size={24} />
      </div>
      <div className={styles.featureContent}>
        <h3 className={styles.featureTitle}>{title}</h3>
        <p className={styles.featureDescription}>{description}</p>
      </div>
    </div>
  );
}

function FeaturesSection() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className={styles.featuresHeader}>
          <Heading as="h2" className={styles.sectionTitle}>
            Key Features
          </Heading>
          <p className={styles.sectionSubtitle}>
            gTunnel offers a range of features designed to enhance your development workflow and collaboration.
          </p>
        </div>
        <div className={styles.featuresGrid}>
          <FeatureCard
            icon={Globe}
            title="Public URLs"
            description="Share your local development server with anyone via a public URL."
          />
          <FeatureCard
            icon={Shield}
            title="Secure Tunnels"
            description="Ensure secure communication with encrypted tunnels."
          />
          <FeatureCard
            icon={Code}
            title="Open Source"
            description="Contribute to and customize gTunnel as an open-source project."
          />
          <FeatureCard
            icon={Zap}
            title="Lightning Fast"
            description="Built with Go for exceptional performance and minimal resource usage."
          />
          <FeatureCard
            icon={Container}
            title="Docker Ready"
            description="Multi-architecture Docker images available for easy deployment."
          />
          <FeatureCard
            icon={Wrench}
            title="Developer Friendly"
            description="Comprehensive CLI tools with detailed documentation and examples."
          />
        </div>
      </div>
    </section>
  );
}

function UseCasesSection() {
  return (
    <section className={styles.useCases}>
      <div className="container">
        <div className={styles.useCasesContent}>
          <div className={styles.useCasesText}>
            <Heading as="h2" className={styles.sectionTitle}>
              Perfect for Development
            </Heading>
            <p className={styles.sectionSubtitle}>
              Common scenarios where gTunnel makes development workflow seamless and efficient.
            </p>
          </div>
          <div className={styles.useCasesList}>
            <div className={styles.useCase}>
              <div className={styles.useCaseIconContainer}>
                <Users className={styles.useCaseIcon} size={20} />
              </div>
              <div className={styles.useCaseContent}>
                <h4>Team Collaboration</h4>
                <p>Share your local development server with teammates instantly</p>
              </div>
            </div>
            <div className={styles.useCase}>
              <div className={styles.useCaseIconContainer}>
                <Webhook className={styles.useCaseIcon} size={20} />
              </div>
              <div className={styles.useCaseContent}>
                <h4>Webhook Testing</h4>
                <p>Test webhooks from external services during development</p>
              </div>
            </div>
            <div className={styles.useCase}>
              <div className={styles.useCaseIconContainer}>
                <Target className={styles.useCaseIcon} size={20} />
              </div>
              <div className={styles.useCaseContent}>
                <h4>Client Demos</h4>
                <p>Showcase your work-in-progress to clients and stakeholders</p>
              </div>
            </div>
            <div className={styles.useCase}>
              <div className={styles.useCaseIconContainer}>
                <FlaskConical className={styles.useCaseIcon} size={20} />
              </div>
              <div className={styles.useCaseContent}>
                <h4>API Integration</h4>
                <p>Debug third-party API integrations in real-time</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

function CTASection() {
  return (
    <section className={styles.cta}>
      <div className="container">
        <div className={styles.ctaContent}>
          <Heading as="h2" className={styles.ctaTitle}>
            Ready to get started?
          </Heading>
          <p className={styles.ctaSubtitle}>
            Join developers worldwide who trust gTunnel for their tunneling needs
          </p>
          <div className={styles.ctaButtons}>
            <Link
              className="button button--primary button--lg"
              to="/quick-start">
              Start Tunneling Now
            </Link>
            <Link
              className="button button--secondary button--lg button--outline"
              href="https://github.com/B-AJ-Amar/gTunnel">
              View on GitHub
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title="Fast & Secure HTTP Tunneling"
      description="gTunnel - Fast, secure, and lightweight HTTP tunneling solution built with Go. Expose your local development servers to the internet instantly.">
      <ScrollNavbar />
      <HomepageHeader />
      <main>
        <FeaturesSection />
        <UseCasesSection />
        <CTASection />
      </main>
    </Layout>
  );
}
