<script lang="ts">
    import type {IProjectItem} from '../../../interfaces/all';
    import ProjectHero from '../../../components/projects/ProjectHero.svelte';

    import {Projects} from '../../../stores/projects';
    import {page} from '$app/stores';
    import Caption from '../../../components/Caption.svelte';
    import Image from '../../../components/Image.svelte';
    import TechLogo from "../../../components/TechLogo.svelte";

    const tokens = $page.url.pathname.split('/');
    const thisProject: IProjectItem = $Projects.find((p) => p.slug.includes(tokens[tokens.length - 1]))
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
        <h2>At a glance</h2>
        <div class="px-3">
            <ul class="list-disc list-outside">
                <li>
                    As the sole developer, designed and implemented an online booking platform and reservation
                    management
                    system.
                </li>
                <li>
                    Successfully launched the platform and achieved a monthly recurring revenue (MRR) of $60,000 in 12
                    months.
                </li>
            </ul>
        </div>
        <h2>What's The Bindel?</h2>
        <p>
            The Bindel is an online rental booking platform for short-term stays in Austin, Texas, offering a fun and
            simple user experience for guests.
        </p>
        <h2>Technology Stack and Implementation Approach</h2>
        <div class="grid fill-current text-gray-700 gap-2 grid-flow-col grid-cols-3 md:grid-cols-6 grid-rows-2 md:grid-rows-1 items-center">
            <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-4 flex flex-col h-full justify-between items-center">
                <div class="w-8 max-w-full text-sm mx-auto">
                    <TechLogo  name="Python 3" />
                </div>
                Python 3
            </div>
            <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex flex-col h-full justify-between items-center">
                <div class="w-12 max-w-full">
                    <TechLogo class="w-8 mx-auto my-auto" name="NuxtJs" />
                </div>
                NuxtJs
            </div>
            <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex flex-col h-full justify-between items-center">
                <div class="w-8 max-w-full">
                    <TechLogo name="Typescript" />
                </div>
                Typescript
            </div>
            <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 h-full flex flex-col justify-center items-center">
                <div class="w-24 max-w-full">
                    <TechLogo name="MongoDb" />
                </div>
            </div>
            <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex h-full flex-col justify-between items-center">
                <div class="w-6 max-w-full">
                    <TechLogo name="Digital Ocean" />
                </div>
                Digital Ocean
            </div>
        </div>
        <p>
            As the sole developer I prioritize choosing technology that allows me to ship features quickly. My go-to
            stack for The Bindel included Python, FastAPI, and MongoDB for the backend, and NuxtJs + TypeScript for the
            frontend. The entire project was hosted on <a href="https://digitalocean.com"
                                                          target="_blank">DigitalOcean</a> and a simple Bash script was
            used for CI/CD.
        </p>
        <p>
            For the frontend, I chose <a href="https://nuxtjs.org/" target="_blank">NuxtJs</a> due to its
            Server-Side-Generation (SSG) capabilities and familiarity with VueJs. Since Nuxt supports both
            SSG for SEO and a fast Single Page Application (SPA) for users, the best of both worlds can be achieved.
            TypeScript's compile-time safety helps get features working with less trial and error.
        </p>
        <p>
            For the backend, I opted for <a href="https://fastapi.tiangolo.com/" target="_blank">FastApi</a> and <a
                href="https://mongodb.com" target="_blank">MongoDB</a>. FastAPI's out-of-the-box OpenAPI support made it
            easy to integrate with the frontend. OpenAPI is also very helpful in testing the backend independently of
            the frontend implementation. MongoDB was a clear choice as it allowed us to move quickly without the
            initial effort of defining a strong data model. Later, Mongo's Aggregation Pipeline enabled sensible
            cross-collection &ldquo;joins&rdquo; after patterns in the data began to emerge.
        </p>
    </div>
</section>

<section>
    <div class="flex flex-row space-x-4 bg-gray-100 rounded p-4 my-2 w-full">
        <div>
            <Image src="/bindel/nav-loc-picker.svg"
                   alt="Wireframe of the Bindel Navigation Menu with Location picker open"/>
        </div>
        <div>
            <Image src="/bindel/nav-unit-picker.svg"
                   alt="Wireframe of the Bindel Navigation Menu with Unit-Type picker open"/>
        </div>
    </div>
    <div>
        <Caption
                props="The Bindel — Wireframes illustrating the location and unit-type pickers in their open and intentionally hap-hazard appearance."/>
    </div>
    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Making it fun</h2>
            <p>
                To make the frontend fun and engaging, despite competition from other players with bigger budgets, we focused on keeping the Bindel site simple and enjoyable. One key area that received extra attention was the neighborhood and unit-type picker. To achieve a seamless and enjoyable experience, the navigation menu haphazardly unfolds into a stack of cards, each featuring a neighborhood with a cute icon to entice interest. Similarly, the unit picker works the same way, with slightly bigger cards to help distinguish amenities.
            </p>
        </div>
    </div>
    <div class="flex space-x-2 flex-row max-w">
        <video width="1200" height="1200" autoplay loop>
            <source src="/bindel/BindelNav.mp4" type="video/mp4">
            <source src="/bindel/BindelNav.mov" type="video/quicktime;codecs=hvec">
        </video>
    </div>


    <div>
        <Caption
                props="Bindel navigation menu implemented."/>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <p>
                For the product detail page (PDP), we relied heavily on previously established patterns, as people already know how to shop stays online. Our efforts focused on making the experience blazing fast. Opening the photo gallery is instantaneous, and so is interacting with the rest of the page, calendar, buy-box, and so on.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Admin Tools</h2>
            <p>
                The website has a significant amount of admin functionality, which includes managing neighborhoods, unit
                types, and pricing, each with its own dedicated admin interface. However, the most critical admin work
                was focused on the guest experience, specifically handling reservations, move-in and move-out,
                invoicing, and payments. To improve upon Stripe's out-of-the-box functionality, we developed our own
                invoicing tool designed to address the intricacies of the rental industry. Our primary goal was to make
                the process simple and error-proof for the customer-facing staff.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Check it out</h2>
            <p>Visit <a target="_blank" href="https://thebindel.com">The Bindel.</a></p>
        </div>
    </div>

</section>

