<script>
    import TheNav from '../components/nav/TheNav.svelte';
    import ContactModal from '../components/ContactModal.svelte';
    import './app.css';
    import {SvelteToast} from '@zerodevx/svelte-toast';
    import TechLogo from "../components/TechLogo.svelte";
</script>

<div class="">
    <TheNav/>
    <ContactModal/>
    <main class="mx-auto max-w-5xl px-4 lg:px-8">
        <slot/>
    </main>

    <footer class="mx-auto mb-10 flex max-w-5xl items-center justify-around pt-6 px-6 lg:px-8 align-middle">
        <p class="text-sm">
            Spence Wetjen is a Full Stack Developer in Austin, Texas.
        </p>
        <a href="https://github.com/swetjen" class="fill-gray-700 p-1 rounded hover:fill-white hover:bg-blue-600" target="_blank">
        <TechLogo class="w-6" name="Github"/>
        </a>
        <p class="text-sm">
            © 2023 Spence Wetjen
        </p>
    </footer>
    <SvelteToast options={{ reversed: true, intro: { y: 192 } }}/>
</div>


<style>
    :root {
        --toastContainerTop: auto;
        --toastContainerRight: auto;
        --toastContainerBottom: 8rem;
        --toastContainerLeft: calc(50vw - 8rem);
    }
</style>

