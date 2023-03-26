<script lang="ts">
  import type {IProjectItem} from '../../interfaces/iprojects';
  export let project: IProjectItem
</script>

<section class="my-10">
    <div class="md:h-96 -z-50 h-64 bg-center bg-cover rounded-2xl shadow-2xl"
         style={`background-image: url('${project.hero_image}')`}>
<!--        <div class="w-full h-full flex flex-row text-center">&nbsp;-->
<!--            <img class="inline-block mx-auto" width="125" src={project.logo}>-->
<!--        </div>-->
    </div>

    <div class="my-10 md:mb-20 flex align-middle md:flex-row mx-5 flex-col">
        <div class="flex-1 sm:w-full md:w-1/2">
            <h1 class="text-3xl font-heavier md:mx-auto">{project.title}</h1>
            <p>September 2020 - June 2022</p>
        </div>
        <p class="text-gray-700">{project.brand}</p>
<!--        <span class="flex-2 no-underline sm:w-full md:w-1/2 my-auto md:mx-auto ">-->
<!--                    {project.summary}-->
<!--        </span>-->
    </div>

</section>