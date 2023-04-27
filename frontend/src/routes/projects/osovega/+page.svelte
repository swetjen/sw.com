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
    <ProjectHero project={thisProject}/>
</section>


<section class="my-10 space-y-8">
    <p>todo</p>
</section>
