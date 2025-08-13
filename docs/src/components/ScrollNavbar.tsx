import { useEffect } from 'react';

export default function ScrollNavbar() {
  useEffect(() => {
    const handleScroll = () => {
      const navbar = document.querySelector('.navbar');
      if (navbar) {
        if (window.scrollY > 50) {
          navbar.classList.add('navbar--scrolled');
          navbar.classList.add('navbar--landing-scrolled');
        } else {
          navbar.classList.remove('navbar--scrolled');
          navbar.classList.remove('navbar--landing-scrolled');
        }
      }
    };

    // Add page-specific classes to navbar and body for landing page
    const navbar = document.querySelector('.navbar');
    const body = document.body;
    
    if (navbar) {
      navbar.classList.add('navbar--landing-transparent');
    }
    
    if (body) {
      body.classList.add('landing-page');
    }

    // Add scroll event listener
    window.addEventListener('scroll', handleScroll);
    
    // Call once to set initial state
    handleScroll();

    // Cleanup
    return () => {
      window.removeEventListener('scroll', handleScroll);
      // Remove page-specific classes when component unmounts
      if (navbar) {
        navbar.classList.remove('navbar--landing-transparent');
        navbar.classList.remove('navbar--landing-scrolled');
        navbar.classList.remove('navbar--scrolled');
      }
      if (body) {
        body.classList.remove('landing-page');
      }
    };
  }, []);

  return null; // This component doesn't render anything
}
