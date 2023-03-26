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
 ]

</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <p>The Bindel</p>
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