<script>
	import TheNav from '../components/nav/TheNav.svelte';
	import ContactModal from '../components/ContactModal.svelte';
	import './app.css';
	import {SvelteToast} from '@zerodevx/svelte-toast';
</script>

<div class="">
	<TheNav />
	<ContactModal />
	<main class="mx-auto max-w-5xl px-4 lg:px-8">
		<slot />
	</main>

	<footer>
		<p class="mt-40 text-center font-bold"></p>
		the footer
	</footer>
	<SvelteToast options={{ reversed: true, intro: { y: 192 } }} />
</div>


<style>
	:root {
		--toastContainerTop: auto;
		--toastContainerRight: auto;
		--toastContainerBottom: 8rem;
		--toastContainerLeft: calc(50vw - 8rem);
	}
</style>

