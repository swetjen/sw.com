<script lang="ts">
  import type {IProjectItem} from '../../../interfaces/iprojects';
  import ProjectHero from '../../../components/projects/ProjectHero.svelte';

  let project: IProjectItem = {
    logo: '/logos/bindel.svg',
    brand: 'The Bindel',
    title: 'Vacation Rentals',
    summary: 'What if you had a flexible housing provider that made it easy to live and work anywhere?',
    hero_image: '/bg/bindel-bg.jpg',
    slug: '/projects/bindel'
  }

</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <ProjectHero project={project}/>
</section>

<section class="my-10 space-y-8">
    <p>
        "As the lead developer for The Bindel's online booking platform, I had the privilege of contributing to the
        success of an Austin, Texas-based startup that provides flexible housing solutions for remote workers. My work
        on the project involved writing both the backend and frontend applications, which included restful APIs, a
        database, and an ERP system.
    </p>
    <p>
        Using the latest technologies such as Python with FastApi for Restful Services, NuxtJS for fast SSG (Server Side
        Generated Pages), and MongoDB for the database, I created a scalable, reliable platform that could handle the
        company's growing customer base. Additionally, I built a frontend interface that enabled customers to see
        apartment availability and book their stay, while also providing an administrative interface for managing
        apartments, tracking availability, reviewing and approving guest reservations, sending invoices, and collecting
        payments.
    <p>
        My work played a significant role in helping The Bindel get in front of customers quickly by iterating rapidly
        and releasing new functionality daily. By building a user-friendly platform that offered seamless and convenient
        housing solutions for remote workers, we were able to contribute to the company's success and help it grow into
        a leading player in the industry."
    </p>
</section>

<section>
    ☐ Better hero image
    ☐ Screenshots
    ☐ Embed the fun menu
    ☐ Snippets for each key screenshot.
    ☐ Fun features section.
    ☐ Technology logos
</section>

<!--<section class="my-10">-->
<!--    {#each projects as p (p.title)}-->
<!--        <a class="my-10" href={p.slug}>-->
<!--            <div class="md:h-96 -z-50 h-64 bg-center bg-cover rounded-2xl transform-gpu hover:scale-95 duration-150 ease-in-out shadow-2xl hover:shadow-3xl "-->
<!--                 style={`background-image: url('${p.hero_image}')`}>-->
<!--                <div class="w-full h-full flex flex-row text-center">&nbsp;-->
<!--                    <img class="inline-block mx-auto" width="125" src={p.logo}>-->
<!--                </div>-->
<!--            </div>-->

<!--            <div class="my-10 md:mb-20 flex align-middle md:flex-row mx-10 flex-col">-->
<!--                <div class="flex-1 sm:w-full md:w-1/2">-->
<!--                    <h1 class="mb-3 text-3xl font-bold md:mx-auto">{p.title}</h1>-->
<!--                </div>-->
<!--                <span class="flex-2 no-underline sm:w-full md:w-1/2 md:mx-auto ">-->
<!--                    {p.summary}-->
<!--                </span>-->
<!--            </div>-->
<!--        </a>-->
<!--    {/each}-->
<!--</section>-->