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
    dates: 'September 2020 - June 2022',
    slug: '/projects/bindel'
  },
  {
    logo: '/logos/tethr.svg',
    brand: 'Tethr',
    title: 'Customer Intelligence',
    titleExtended: 'Enterprise scale call-center conversations analyzed, automated and audited',
    // summary: 'Tethr helps enterprise scale companies to sift through millions of customer phone calls surfacing latent customer insights, improving representative performance, and ensure compliance.',
    heroImage: '/bg/tethr.webp',
    dates: '2012 - 2014',
    slug: '/projects/tethr',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Digital Tools',
    brand: 'Best Buy',
    titleExtended: 'Bringing efficiency to a dozens of adhoc customer and employee applications',
    heroImage: '/bg/digitaltools.webp',
    dates: '2010 - 2012',
    slug: '/projects/bestbuy-digital-tools',
  },
  {
    logo: '/logos/bby.svg',
    title: 'Store Demos',
    brand: 'Best Buy',
    titleExtended: 'Instore demo for laptop purchasers',
    heroImage: '/bg/bbi.webp',
    dates: '2008 - 2009',
    slug: '/projects/bestbuy-store-demo',
  },
  {
    logo: '/logos/ea.svg',
    title: 'Retail Kiosks',
    brand: 'Best Buy',
    titleExtended: 'Browsable and pre-order upcoming video games',
    heroImage: '/bg/pots.webp',
    dates: '2009 - 2010',
    slug: '/projects/ea-gaming-kiosk',
  },
]

export const Projects: Readable<IProjectItem[]> = readable(projectsDefault)

