import {readable} from 'svelte/store';
import type { Readable } from 'svelte/store';
import type { IProjectItem } from '../interfaces/all';


let projectsDefault: IProjectItem[] = [
  {
    logo: '/logos/bindel.svg',
    brand: 'The Bindel',
    heroImage: '/bg/bindel-bg.jpg',
    title: 'Reservations Platform',
    titleExtended: 'An online booking platform that let guests browse, book and pay',
    slug: '/projects/bindel'
  },
  {
    logo: '/logos/tethr.svg',
    brand: 'Tethr',
    title: 'Customer Intelligence',
    titleExtended: 'What if customer and agent conversations could be searched, analyzed and automated?',
    // summary: 'Tethr helps enterprise scale companies to sift through millions of customer phone calls surfacing latent customer insights, improving representative performance, and ensure compliance.',
    heroImage: '/bg/tethr.webp',
    slug: '/projects/tethr',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Digital Tools',
    brand: 'Best Buy',
    titleExtended: 'How do we bring order and efficiency to a dozens of one-off customer and employee applications?',
    heroImage: '/bg/digitaltools.webp',
    slug: '/projects/bestbuy-digital-tools',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Store Demos',
    brand: 'Best Buy',
    titleExtended: 'Building tools to help prospective laptop buyers compare product features.',
    heroImage: '/bg/bbi.webp',
    slug: '/projects/bestbuy-store-demo',
  },
  {
    logo: '/logos/ea.svg',
    title: 'Retail Kiosks',
    brand: 'Best Buy',
    titleExtended: 'What if we build a browsable catalog of coming soon games?',
    heroImage: '/bg/pots.webp',
    slug: '/projects/ea-gaming-kiosk',
  },
]

export const Projects: Readable<IProjectItem[]> = readable(projectsDefault)

