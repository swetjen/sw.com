<script lang="ts">
  import { fade, fly } from 'svelte/transition';
  import { page } from '$app/stores';
  import {Contact} from '../../stores/stores';
  import type {IProjectItem} from '../../interfaces/all';
  import { Projects } from '../../stores/projects';
  import NavProject from './NavProject.svelte';

  // menu and contact states
  let ProductOpen: boolean = false
  let MainMenuOpen: boolean = false

  function Reset() {
    ProductOpen = false
    MainMenuOpen = false
  }

</script>

<!--<header>-->
<!--    <nav>-->
<!--        <div>-->
<!--            <div class="mx-auto flex justify-between items-center space-x-6 mt-4">-->
<!--                <a class="flex-1 text-4xl font-bold" href="/">-->
<!--                    Spence Wetjen-->
<!--                </a>-->
<!--&lt;!&ndash;                <a aria-current={$page.url.pathname === '/' ? 'page' : undefined} href="/">Home</a>&ndash;&gt;-->
<!--                <a aria-current={$page.url.pathname === '/' ? 'page' : undefined} href="/">Projects</a>-->
<!--                <a aria-current={$page.url.pathname === '/' ? 'page' : undefined} href="/about">About</a>-->
<!--                <a href="/">Contact</a>-->
<!--            </div>-->
<!--        </div>-->
<!--    </nav>-->
<!--</header>-->
{#if ProductOpen || MainMenuOpen }
    <!--
        Todo - Transition
    -->
    <div transition:fade="{{ duration: 200 }}"  on:click={Reset} class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
{/if}

<!--<button on:click={() => MainMenuOpen = !MainMenuOpen} type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700">-->


<header class="relative isolate z-10 bg-white">
    <nav class="mx-auto flex max-w-5xl items-center justify-between pt-6 lg:px-8" aria-label="Global">
        <div class="flex lg:flex-1">
            <a href="/" class="-m-1.5 p-1.5">
                <h1 aria-description="Spence Wetjen">Spence Wetjen</h1>
            </a>
        </div>
        <div class="flex lg:hidden">
            <button type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700">
                <span class="sr-only">Open main menu</span>
                <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                </svg>
            </button>
        </div>
        <div class="hidden lg:flex lg:gap-x-12">
            <div>
                <button on:click={() => ProductOpen = !ProductOpen} type="button" class="flex items-center gap-x-1 nav-link leading-6 " aria-expanded="false">
                    Work
                    <svg class="h-5 w-5 flex-none text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                        <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
                    </svg>
                </button>

                <!--
                  'Product' flyout menu, show/hide based on flyout menu state.

                  Entering: "transition ease-out duration-200"
                    From: "opacity-0 -translate-y-1"
                    To: "opacity-100 translate-y-0"
                  Leaving: "transition ease-in duration-150"
                    From: "opacity-100 translate-y-0"
                    To: "opacity-0 -translate-y-1"
                -->
                {#if ProductOpen}
                <div class="absolute  inset-x-0 top-0 -z-10 bg-white pt-14 shadow-lg ring-1 ring-gray-900/5">
                    <div class="mx-auto grid max-w-5xl grid-cols-3 gap-4 px-6 py-10 lg:px-8 xl:gap-x-8">

                        {#each $Projects as p (p.title)}
                            <NavProject props={p} on:click={Reset} />
                        {/each}

                    </div>
                </div>
                {/if}
            </div>

            <a href="/about" class:active={$page.url.pathname === "/about"} class="leading-6 nav-link">About</a>
        </div>
        <div class="hidden lg:flex lg:flex-1 lg:justify-end">
            <button on:click={Contact.toggle} class="leading-6 text-2xl hover:bg-blue-600 font-light text-white flex-4 bg-blue-800 rounded px-3 py-2 hidden md:inline-block">Contact &rarr;<span aria-hidden="true"></span></button>
        </div>
    </nav>
    <!-- Mobile menu, show/hide based on menu open state. -->
    <div class="lg:hidden" role="dialog" aria-modal="true">
        <!-- Background backdrop, show/hide based on slide-over state. -->
        <div class="fixed inset-0 z-10"></div>
        <div class="fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
            <div class="flex items-center justify-between">
                <a href="#" class="-m-1.5 p-1.5">
                    <span class="sr-only">Your Company</span>
                    <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="">
                </a>
                <button type="button" class="-m-2.5 rounded-md p-2.5 text-gray-700">
                    <span class="sr-only">Close menu</span>
                    <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>
            <div class="mt-6 flow-root">
                <div class="-my-6 divide-y divide-gray-500/10">
                    <div class="space-y-2 py-6">
                        <div class="-mx-3">
                            <button type="button" class="flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 hover:bg-gray-50" aria-controls="disclosure-1" aria-expanded="false">
                                Product
                                <!--
                                  Expand/collapse icon, toggle classes based on menu open state.

                                  Open: "rotate-180", Closed: ""
                                -->
                                <svg class="h-5 w-5 flex-none" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                                    <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
                                </svg>
                            </button>
                            <!-- 'Product' sub-menu, show/hide based on menu state. -->
                            <div class="mt-2 space-y-2" id="disclosure-1">
                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Analytics</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Engagement</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Security</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Integrations</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Watch demo</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">Contact sales</a>

                                <a href="#" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">View all products</a>
                            </div>
                        </div>
                        <a href="#" class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Features</a>
                        <a href="#" class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Marketplace</a>
                        <a href="#" class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Company</a>
                    </div>
                    <div class="py-6">
                        <a href="#" class="-mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Log in</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</header>