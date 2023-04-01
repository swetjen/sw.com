<script lang="ts">
  import type {IProjectItem} from '../../../interfaces/all';
  import ProjectHero from '../../../components/projects/ProjectHero.svelte';
  import TheBread from '../../../components/TheBread.svelte';

  import {Projects} from '../../../stores/projects';
  import {page} from '$app/stores';
  const tokens = $page.url.pathname.split('/');
  const thisProject: IProjectItem = $Projects.find((p) => p.slug.includes(tokens[tokens.length -1]))


</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <TheBread/>
    <ProjectHero project={thisProject}/>
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

<section class="my-10 space-y-8">
    <p>todo</p>
</section>
