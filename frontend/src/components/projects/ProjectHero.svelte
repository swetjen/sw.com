<script lang="ts">
  import type {IProjectItem} from '../../interfaces/all';
  export let project: IProjectItem
</script>

<section class="max-w-2xl mt-10 mx-auto">
    <div class="my-10 flex align-middle md:flex-row mx-5 flex-col">
        <div class="flex-1 sm:w-full md:max-w-lg">
            <p class="text-gray-500 my-5 text-sm leading-6">{project.dates}/p>
            <h1>{project.titleExtended}</h1>
            {#if project.subtitle}
            <h2 class="my-5">{project.subtitle}</h2>
            {/if}
        </div>
        <!--        <span class="flex-2 no-underline sm:w-full md:w-1/2 my-auto md:mx-auto ">-->
        <!--                    {project.summary}-->
        <!--        </span>-->
    </div>
</section>

<section class="my-10">
    <div class="md:h-96 -z-50 h-64 bg-center bg-cover rounded-2xl shadow-2xl"
         style={`background-image: url('${project.heroImage}')`}>
<!--        <div class="w-full h-full flex flex-row text-center">&nbsp;-->
<!--            <img class="inline-block mx-auto" width="125" src={project.logo}>-->
<!--        </div>-->
    </div>
</section>