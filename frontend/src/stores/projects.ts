import {readable} from 'svelte/store';
import type { Readable } from 'svelte/store';
import type { IProjectItem } from '../interfaces/all';


const projectsDefault: IProjectItem[] = [
  {
    logo: '/logos/bindel.svg',
    brand: 'The Bindel',
    heroImage: '/bg/bindel-bg.jpg',
    title: 'Custom Rental Reservations Platform',
    titleExtended: 'An online booking platform that let guests browse, book and pay',
    subtitle: 'An online rental reservations and booking platform.',
    dates: '2020 - 2022',
    slug: '/projects/bindel'
  },
  {
    logo: '/logos/ov-logo-white.svg',
    brand: 'Osovega',
    title: 'Trading Platform for Specialty Crops',
    subtitle: 'Helping buyers and growers meet and transact with confidence.',
    titleExtended: 'Helping buyers and growers meet and transact with confidence.',
    heroImage: '/bg/ov-hero-bg.jpg',
    dates: '2018 - 2020',
    slug: '/projects/osovega',
  },
  {
    logo: '/logos/tethr.svg',
    brand: 'Tethr',
    title: 'AI Conversation Intelligence',
    titleExtended: 'Tethr is a platform that uses AI-powered conversation intelligence to help businesses improve customer interactions and drive better outcomes.',
    subtitle: 'Conversation intelligence to help businesses improve customer interactions.',
    heroImage: '/bg/tethr.webp',
    dates: '2012 - 2014',
    slug: '/projects/tethr',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Digital Tools',
    subtitle: 'Best Buy\'s Journey to a Cohesive Digital Strategy',
    brand: 'Best Buy',
    titleExtended: 'Aligning Apps and Tools with Cohesive Team and Application Integration',
    heroImage: '/bg/digitaltools.webp',
    dates: '2010 - 2012',
    slug: '/projects/bestbuy-digital-tools',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Store Demos',
    brand: 'Best Buy',
    titleExtended: 'Instore demo for laptop shoppers',
    subtitle: 'Coming soon.',
    heroImage: '/bg/bbi.webp',
    dates: '2008 - 2009',
    slug: '/projects/bestbuy-store-demo',
  },
  {
    logo: '/logos/ea.svg',
    title: 'Retail Kiosks',
    brand: 'Best Buy',
    titleExtended: 'Browsable and pre-order upcoming video games',
    subtitle: 'Coming soon.',
    heroImage: '/bg/pots.webp',
    dates: '2009 - 2010',
    slug: '/projects/ea-gaming-kiosk',
  },
]

export const Projects: Readable<IProjectItem[]> = readable(projectsDefault)

// Click Clinician