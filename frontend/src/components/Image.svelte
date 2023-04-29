<script lang="ts">
  import { onMount } from 'svelte';
  import {fade, fly} from 'svelte/transition';

  export let src: string
  export let fullSrc: string
  export let alt: string
  export let fullImgHint = false

  let ImageExpanded = false

  let y;

  onMount(async () => {
      if (typeof fullSrc === 'undefined') {
          fullSrc = src
      }
  });

</script>

<div class={$$props.class}>
    <img class="cursor-zoom-in focus:ring-2 hover:ring-2 rounded ring-blue-500" on:click={() => ImageExpanded = !ImageExpanded} src="{src}" alt="{alt}">
</div>

<svelte:window bind:scrollY={y}/>

{#if ImageExpanded}
    <div>
        <div class="fixed z-10 inset-0 overflow-y-auto">
            <div class="flex md:items-end justify-center pt-4 px-4 text-center sm:block sm:p-0">

                <!--
              Background overlay, show/hide based on modal state.

              Entering: "ease-out duration-300"
                From: "opacity-0"
                To: "opacity-100"
              Leaving: "ease-in duration-200"
                From: "opacity-100"
                To: "opacity-0"
            -->
                <div in:fade="{{ duration: 300}}" out:fade="{{duration: 200}}" class="fixed inset-0 transition-opacity"
                     aria-hidden="true">
                    <div
                            on:click={() => ImageExpanded = !ImageExpanded}
                            class="absolute inset-0 bg-gray-900 opacity-75"></div>
                </div>

                <!-- This element is to trick the browser into centering the modal contents. -->
<!--                <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>-->
                <!--
                -->
                <div in:fly="{{ duration: 300, y: 18}}" out:fly="{{duration:200, y:18}}"
                     class="inline-block z-20 align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle"
                     role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                         <img class="cursor-zoom-out" src="{fullSrc}" alt="{alt}" on:click={() => ImageExpanded = !ImageExpanded}>

                </div>
            </div>
        </div>
    </div>
{/if}
