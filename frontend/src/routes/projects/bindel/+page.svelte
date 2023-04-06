<script lang="ts">
  import type {IProjectItem} from '../../../interfaces/all';
  import ProjectHero from '../../../components/projects/ProjectHero.svelte';
  import TheBread from '../../../components/TheBread.svelte';

  import {Projects} from '../../../stores/projects';
  import {page} from '$app/stores';
  import Caption from '../../../components/Caption.svelte';
  const tokens = $page.url.pathname.split('/');
  const thisProject: IProjectItem = $Projects.find((p) => p.slug.includes(tokens[tokens.length -1]))
</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <ProjectHero project={thisProject}/>
</section>

<section class="my-10 max-w-2xl mx-auto">
    <div class="mx-5 space-y-8">
    <p>
        "As the lead developer for The Bindel's online booking platform, I had the privilege of contributing to the
        success of an Austin, Texas-based startup that provides flexible housing solutions for remote workers. My work
        on the project involved writing both the backend and frontend applications, which included restful APIs, a
        database, and an ERP system.
    </p>
        <div class="flex space-x-2 flex-row">
        <video width="300" height="300" autoplay loop>
            <source src="/videos/BindelLocationPicker.mp4" type="video/mp4">
        </video>
        </div>
    <h2>Making it fun</h2>
    <p>
    Using the latest technologies including Python with FastApi for Restful Services, NuxtJS for fast SSG (Server Side
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
    </div>
    <div class="flex flex-row space-x-2 bg-gray-300 rounded p-4 my-6 w-full">
        <div>
            <img src="/bindel/pdp.svg"/>
        </div>
        <div>
            <img src="/bindel/pdp-upgrade-selected.svg"/>
        </div>
    </div>
</section>

<section>
    <div class="flex flex-row space-x-2 bg-gray-300 rounded p-4 my-2 w-full">
        <div>
            <img src="/bindel/nav-loc-picker.svg"/>
        </div>
        <div>
            <img src="/bindel/nav-unit-picker.svg"/>
        </div>
    </div>
    <div>
        <Caption props="The Bindel — Wireframes illustrating the location and unit-type pickers in their open and intentionally hap-hazard appearance."/>
    </div>

</section>

<section>
    Notion of making it simple, fun.
    Happy path is great; but upgrades make it amazing.
    ☐ Better hero image
    ☐ Screenshots
    ☐ Embed the fun menu
    ☐ Snippets for each key screenshot.
    ☐ Fun features section.
    ☐ Technology logos
</section>
