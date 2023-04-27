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
    <ul class="list-disc list-outside">
        <li>Co-founded the company.</li>
        <li>Designed an award-winning user experience (UX) that facilitates easy navigation of conversations.</li>
        <li>Successfully launched the platform and achieved a monthly recurring revenue (MRR) of $60,000 in 12 months.
        </li>
        <li>Led a team of 12 engineers by recruiting, managing and guiding them towards success.</li>
    </ul>
    <h2>The Problem</h2>
    <p>
        After spending a decade in Fortune 50 corporate America, I personally experienced the immense size and
        complexity of large organizations. Ensuring that all employees are aligned with the company's current messaging
        can be incredibly challenging, and certain types of messaging can carry significant risks if not executed
        correctly, such as legal and compliance issues.
    <p>Tethr was born out of the idea of simplifying things by automatically documenting employee-customer interactions.
    Since its automatic it wouldn't rely on note-taking or the employee's self-awareness. Aggregated and used a learning tool, it would become a virtuous cycle leading to on-message employees and better outcomes for customers.
    </p>
    <h2>Concept Development</h2>
    <p>
        I worked through several concepts before finalizing the double wave-form inspired UI, which subtly indicates to
        users
        that they are working with audio. The callouts along the wave-form highlight key and make the flow of the
        conversation visible at a glance.
        The component is inherently mobile-friendly, works for conversations of most lengths, and allows for easy
        navigation of the conversation, much
        like a video editing tool.
    </p>
</section>
<div class="flex space-x-2 flex-row max-w">
    <video width="1200" height="1200" autoplay loop>
        <source src="/videos/alt.mp4" type="video/mp4">
    </video>
</div>
<section class="my-10 max-w-2xl mx-auto space-y-8">
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
