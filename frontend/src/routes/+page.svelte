<script lang="ts">
  interface IProjectItem {
    logo: string
    brand: string
    hero_image: string
    title: string
    summary: string
    slug: string
  }

  let projects: IProjectItem[] = [
    {
      logo: '/logos/bindel.svg',
      brand: 'The Bindel',
      title: 'Vacation Rentals',
      summary: 'What if you had a flexible housing provider that made it easy to live and work anywhere?',
      hero_image: '/bg/bindel-bg.jpg',
      slug: '/projects/bindel'
    },
    {
      logo: '/logos/tethr.svg',
      brand: 'Tethr',
      title: 'Conversation Intelligence',
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
      title: 'Store Demo',
      brand: 'Best Buy',
      summary: 'Building tools to help prospective laptop buyers compare product features.',
      hero_image: '/bg/bbi.webp',
      slug: '/projects/best-buy-interactive',
    },
    {
      logo: '/logos/ea.svg',
      title: 'Retail Kiosk',
      brand: 'Best Buy',
      summary: 'What if we build a browsable catalog of coming soon games?',
      hero_image: '/bg/pots.webp',
      slug: '/projects/gaming-kiosk',
    },
  ]

</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <h1 class="text-5xl">Todo - Intro text goes here</h1>
</section>

<section class="my-10">
    {#each projects as p (p.title)}
        <a class="my-10" href={p.slug}>
            <div class="md:h-96 -z-50 h-64 bg-center bg-cover rounded-2xl transform-gpu hover:scale-95 duration-150 ease-in-out shadow-2xl hover:shadow-3xl "
                 style={`background-image: url('${p.hero_image}')`}>
                <div class="w-full h-full flex flex-row text-center">&nbsp;
                    <img class="inline-block mx-auto" width="125" src={p.logo}>
                </div>
            </div>

            <div class="my-10 md:mb-20 flex align-middle md:flex-row mx-10 flex-col">
                <div class="flex-1 sm:w-full md:w-1/2">
                    <h1 class="mb-3 text-3xl font-bold md:mx-auto">{p.title}</h1>
                </div>
                <span class="flex-2 no-underline sm:w-full md:w-1/2 md:mx-auto ">
                    {p.summary}
                </span>
            </div>
        </a>
    {/each}
</section>