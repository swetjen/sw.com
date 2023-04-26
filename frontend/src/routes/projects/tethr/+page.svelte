<script lang="ts">
  import type {IProjectItem} from '../../../interfaces/all';
  import ProjectHero from '../../../components/projects/ProjectHero.svelte';

  import {Projects} from '../../../stores/projects';
  import {page} from '$app/stores';

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

<section class="my-10 max-w-2xl mx-auto space-y-8">
    <h2>At a glance</h2>
    <ul class="list-disc">
        <li>Co-founded the company.</li>
        <li>Designed an award-winning user experience (UX) that facilitates easy navigation of conversations.</li>
        <li>Successfully launched the platform and achieved a monthly recurring revenue (MRR) of $60,000 in 12 months.
        </li>
        <li>Led a team of 12 engineers by recruiting, managing and guiding them towards success.</li>
    </ul>
    <h2>The Problem</h2>
    <p>
        Having spent a decade in Fortune 50 corporate America, I witnessed firsthand the sheer size and complexity of
        large organizations. Aligning staff with the latest brand positioning and standard operating procedures (SOP)
        can be incredibly challenging. </p>
    <p>Tethr was born out of the idea of simplifying things automatically documenting employee-customer interactions.
        This virtuous cycle could lead to better outcomes for customers and more effective employees, without relying on
        note-taking or the employee's self-awareness.
    </p>
    <h2>Concept Development</h2>
    <p>
        I worked through several concepts before finalizing the wave-form inspired UI, which subtly indicates to users
        that they are working with audio. The callouts and double wave show what is happening and highlight key
        concerns. The component is inherently mobile-friendly and allows for easy navigation of the conversation, much
        like a video editing tool.
    </p>
    <p>
        Using this detail view as a basis, views on top of the UI becomes effortless. Begin by creating a persona or
        role interested in specific things, which then become a filterable expression shown in a standard business
        intelligence style view. This seamless process enables executives to extract conversation intelligence without
        needing to use Tethr.
    </p>
    <p>
        Unlike many startups from the that time, Tethr is still alive and growing.
    </p>
    <h2>Check it out</h2>
    <p>Visit <a target="_blank" href="https://tethr.com">Tethr</a>.</p>
</section>
