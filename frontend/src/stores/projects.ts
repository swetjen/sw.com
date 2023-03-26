import {readable} from 'svelte/store';
import type { Readable } from 'svelte/store';
import type { IProjectItem } from '../interfaces/iprojects';

let projectsDefault: IProjectItem[] = [
  {
    logo: '/logos/bindel.svg',
    brand: 'The Bindel',
    title: 'Reservations Platform',
    summary: 'What if you had a flexible housing provider that made it easy to live and work anywhere?',
    hero_image: '/bg/bindel-bg.jpg',
    slug: '/projects/bindel'
  },
  {
    logo: '/logos/tethr.svg',
    brand: 'Tethr',
    title: 'Customer Intelligence',
    summary: 'What if customer and agent conversations could be searched, analyzed and automated?',
    // summary: 'Tethr helps enterprise scale companies to sift through millions of customer phone calls surfacing latent customer insights, improving representative performance, and ensure compliance.',
    hero_image: '/bg/tethr.webp',
    slug: '/projects/tethr',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Digital Tools',
    brand: 'Best Buy',
    summary: 'How do we bring order and efficiency to a dozens of one-off customer and employee applications?',
    hero_image: '/bg/digitaltools.webp',
    slug: '/projects/digital-tools',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Store Demos',
    brand: 'Best Buy',
    summary: 'Building tools to help prospective laptop buyers compare product features.',
    hero_image: '/bg/bbi.webp',
    slug: '/projects/best-buy-interactive',
  },
  {
    logo: '/logos/ea.svg',
    title: 'Retail Kiosks',
    brand: 'Best Buy',
    summary: 'What if we build a browsable catalog of coming soon games?',
    hero_image: '/bg/pots.webp',
    slug: '/projects/gaming-kiosk',
  },
]

export const Projects: Readable<IProjectItem[]> = readable(projectsDefault)

