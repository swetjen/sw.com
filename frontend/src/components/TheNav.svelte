<script lang="ts">
  import { fade, fly, blur } from 'svelte/transition';
  import type {IProjectItem} from '../interfaces/iprojects';

  // Projects data
  import {Projects} from '../stores/projects';

  let projects: IProjectItem[]

  Projects.subscribe(value => {
    projects = value
  })

  // menu and contact states
  let ProductOpen: boolean = false
  let MainMenuOpen: boolean = false
  let ContactModalOpen: boolean = false

  // convience methods
  function Reset() {
    ProductOpen = false
    MainMenuOpen = false
    ContactModalOpen = false
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

<!-- Gray Overlay -->
{#if ProductOpen || MainMenuOpen || ContactModalOpen}
    <!--
        Todo - Transition
    -->
    <div transition:fade="{{ duration: 200 }}"  on:click={Reset} class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
{/if}

<header class="bg-white z-10">
    <nav class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
        <div class="flex lg:flex-1">
            <a href="/" class="-m-1.5 p-1.5 text-4xl font-bold">
                Spence Wetjen
            </a>
        </div>
        <div class="flex lg:hidden">
            <button on:click={() => MainMenuOpen = !MainMenuOpen} type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700">
                <span class="sr-only">Open main menu</span>
                <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                </svg>
            </button>
        </div>
        <div class="hidden lg:flex lg:gap-x-12">
            <div class="relative">
                <button on:click={() => ProductOpen = !ProductOpen} type="button" class="flex items-center gap-x-1 leading-6 text-gray-900" aria-expanded="false">
                    Products
                    <svg class="h-5 w-5 flex-none text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                        <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
                    </svg>
                </button>

                <!--
                  'Product' flyout menu, show/hide based on flyout menu state.

                  Entering: "transition ease-out duration-200"
                    From: "opacity-0 translate-y-1"
                    To: "opacity-100 translate-y-0"
                  Leaving: "transition ease-in duration-150"
                    From: "opacity-100 translate-y-0"
                    To: "opacity-0 translate-y-1"
                -->
                {#if ProductOpen}
                <div transition:fly="{{ y: 20, duration: 100 }}" class="absolute inset-0 z-10 -ml-4 px-2 w-full sm:px-0 lg:ml-0 lg:left-1/2 lg:-translate-x-1/2">
                    <div class="relative bg-transparent grid grid-rows-2 gap-x-12 grid-cols-3 grid-flow-row px-5 z-10 sm:gap-y-4 sm:p-4">
                        {#each projects as p, i}
                            <div
                                class="first:-rotate-2 [&:nth-child(3)]:rotate-2 [&:nth-child(2)]:z-20  col-span-1 bg-grey-10 z-10 hover:z-30 transition shadow-2xl ease-in-out hover:rotate-0 bg-white  transform hover:scale-105 rounded-xl group hover:bg-blue-200 overflow-hidden w-64 shadow-2xl">

                                <!-- Heroicon name: chart-bar -->
                                <div class="bg-cover bg-center h-48 bg-no-repeat" style={`background-image: url('${p.hero_image}')`}>
                                </div>
                                <div
                                        class="font-medium z-50 my-1 mx-3 text-lg text-left text-black group-hover:text-grey-500">
                                    { p.brand }
                                    <div class="text-sm -mt-1 group-hover:text-grey-10 text-grey-200">
                                        { p.title }
                                    </div>
                                    <div class="mt-1 group-hover:text-grey-10 text-grey-100">
                                        <small class="text-xs">Starting at</small>
                                    </div>
                                    <div class="font-medium  -mt-1 text-left text-blue-200 group-hover:text-grey-10">
                                        last row
                                    </div>
                                </div>
                            </div>
                        {/each}



                        <!-- Original Menu Item -->
<!--                        <div class="group relative flex gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">-->
<!--                            <div class="mt-1 flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">-->
<!--                                <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">-->
<!--                                    <path stroke-linecap="round" stroke-linejoin="round" d="M10.5 6a7.5 7.5 0 107.5 7.5h-7.5V6z" />-->
<!--                                    <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 10.5H21A7.5 7.5 0 0013.5 3v7.5z" />-->
<!--                                </svg>-->
<!--                            </div>-->
<!--                            <div class="flex-auto">-->
<!--                                <a href="#" class="block font-semibold text-gray-900">-->
<!--                                    Analytics-->
<!--                                    <span class="absolute inset-0"></span>-->
<!--                                </a>-->
<!--                                <p class="mt-1 text-gray-600">Get a better understanding of your traffic</p>-->
<!--                            </div>-->
<!--                        </div>-->

                    </div>
                </div>
                {/if}
            </div>

            <a href="#" class="leading-6 text-gray-900">Features</a>
            <a href="#" class="leading-6 text-gray-900">Marketplace</a>

<!--            <div class="relative">-->
<!--                <button type="button" class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900" aria-expanded="false">-->
<!--                    Company-->
<!--                    <svg class="h-5 w-5 flex-none text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">-->
<!--                        <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />-->
<!--                    </svg>-->
<!--                </button>-->

<!--                &lt;!&ndash;-->
<!--                  'Company' flyout menu, show/hide based on flyout menu state.-->

<!--                  Entering: "transition ease-out duration-200"-->
<!--                    From: "opacity-0 translate-y-1"-->
<!--                    To: "opacity-100 translate-y-0"-->
<!--                  Leaving: "transition ease-in duration-150"-->
<!--                    From: "opacity-100 translate-y-0"-->
<!--                    To: "opacity-0 translate-y-1"-->
<!--                &ndash;&gt;-->
<!--                <div class="absolute -left-8 top-full z-10 mt-3 w-96 rounded-3xl bg-white p-4 shadow-lg ring-1 ring-gray-900/5">-->
<!--                    <div class="relative rounded-lg p-4 hover:bg-gray-50">-->
<!--                        <a href="#" class="block text-sm font-semibold leading-6 text-gray-900">-->
<!--                            About us-->
<!--                            <span class="absolute inset-0"></span>-->
<!--                        </a>-->
<!--                        <p class="mt-1 text-sm leading-6 text-gray-600">Learn more about our company values and mission to empower others</p>-->
<!--                    </div>-->

<!--                    <div class="relative rounded-lg p-4 hover:bg-gray-50">-->
<!--                        <a href="#" class="block text-sm font-semibold leading-6 text-gray-900">-->
<!--                            Careers-->
<!--                            <span class="absolute inset-0"></span>-->
<!--                        </a>-->
<!--                        <p class="mt-1 text-sm leading-6 text-gray-600">Looking for you next career opportunity? See all of our open positions</p>-->
<!--                    </div>-->

<!--                    <div class="relative rounded-lg p-4 hover:bg-gray-50">-->
<!--                        <a href="#" class="block text-sm font-semibold leading-6 text-gray-900">-->
<!--                            Support-->
<!--                            <span class="absolute inset-0"></span>-->
<!--                        </a>-->
<!--                        <p class="mt-1 text-sm leading-6 text-gray-600">Get in touch with our dedicated support team or reach out on our community forums</p>-->
<!--                    </div>-->

<!--                    <div class="relative rounded-lg p-4 hover:bg-gray-50">-->
<!--                        <a href="#" class="block text-sm font-semibold leading-6 text-gray-900">-->
<!--                            Blog-->
<!--                            <span class="absolute inset-0"></span>-->
<!--                        </a>-->
<!--                        <p class="mt-1 text-sm leading-6 text-gray-600">Read our latest announcements and get perspectives from our team</p>-->
<!--                    </div>-->
<!--                </div>-->
<!--            </div>-->
        </div>
        <div class="hidden lg:flex lg:flex-1 lg:justify-end">
            <a href="#" class="leading-6 text-gray-900">Contact me<span aria-hidden="true"></span></a>
        </div>
    </nav>
    <!-- Mobile menu, show/hide based on menu open state. -->
    <div class="lg:hidden" role="dialog" aria-modal="true">
        <!-- Background backdrop, show/hide based on slide-over state. -->
        <div class="fixed inset-0 z-10"></div>
        <div class="fixed inset-y-0 right-0 z-10 flex w-full flex-col justify-between overflow-y-auto bg-white sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
            <div class="p-6">
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
                            <a href="#" class="group -mx-3 flex items-center gap-x-6 rounded-lg p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M10.5 6a7.5 7.5 0 107.5 7.5h-7.5V6z" />
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 10.5H21A7.5 7.5 0 0013.5 3v7.5z" />
                                    </svg>
                                </div>
                                Analytics
                            </a>

                            <a href="#" class="group -mx-3 flex items-center gap-x-6 rounded-lg p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.042 21.672L13.684 16.6m0 0l-2.51 2.225.569-9.47 5.227 7.917-3.286-.672zM12 2.25V4.5m5.834.166l-1.591 1.591M20.25 10.5H18M7.757 14.743l-1.59 1.59M6 10.5H3.75m4.007-4.243l-1.59-1.59" />
                                    </svg>
                                </div>
                                Engagement
                            </a>

                            <a href="#" class="group -mx-3 flex items-center gap-x-6 rounded-lg p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M7.864 4.243A7.5 7.5 0 0119.5 10.5c0 2.92-.556 5.709-1.568 8.268M5.742 6.364A7.465 7.465 0 004.5 10.5a7.464 7.464 0 01-1.15 3.993m1.989 3.559A11.209 11.209 0 008.25 10.5a3.75 3.75 0 117.5 0c0 .527-.021 1.049-.064 1.565M12 10.5a14.94 14.94 0 01-3.6 9.75m6.633-4.596a18.666 18.666 0 01-2.485 5.33" />
                                    </svg>
                                </div>
                                Security
                            </a>

                            <a href="#" class="group -mx-3 flex items-center gap-x-6 rounded-lg p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 16.875h3.375m0 0h3.375m-3.375 0V13.5m0 3.375v3.375M6 10.5h2.25a2.25 2.25 0 002.25-2.25V6a2.25 2.25 0 00-2.25-2.25H6A2.25 2.25 0 003.75 6v2.25A2.25 2.25 0 006 10.5zm0 9.75h2.25A2.25 2.25 0 0010.5 18v-2.25a2.25 2.25 0 00-2.25-2.25H6a2.25 2.25 0 00-2.25 2.25V18A2.25 2.25 0 006 20.25zm9.75-9.75H18a2.25 2.25 0 002.25-2.25V6A2.25 2.25 0 0018 3.75h-2.25A2.25 2.25 0 0013.5 6v2.25a2.25 2.25 0 002.25 2.25z" />
                                    </svg>
                                </div>
                                Integrations
                            </a>

                            <a href="#" class="group -mx-3 flex items-center gap-x-6 rounded-lg p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <svg class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />
                                    </svg>
                                </div>
                                Automations
                            </a>
                        </div>
                        <div class="space-y-2 py-6">
                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Features</a>
                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Marketplace</a>

                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">About us</a>

                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Careers</a>

                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Support</a>

                            <a href="#" class="-mx-3 block rounded-lg py-2 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Blog</a>
                        </div>
                        <div class="py-6">
                            <a href="#" class="-mx-3 block rounded-lg py-2.5 px-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">Log in</a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="sticky bottom-0 grid grid-cols-2 divide-x divide-gray-900/5 bg-gray-50 text-center">
                <a href="#" class="p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-100">Watch demo</a>

                <a href="#" class="p-3 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-100">Contact sales</a>
            </div>
        </div>
    </div>
</header>
