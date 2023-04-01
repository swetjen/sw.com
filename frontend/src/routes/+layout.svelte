<script>
	import TheNav from '../components/nav/TheNav.svelte';
	import ContactModal from '../components/ContactModal.svelte';
	import './app.css';
</script>

<div class="">
	<TheNav />
	<ContactModal />

	<main class="mx-auto max-w-5xl px-4 lg:px-8">
		<slot />
	</main>

	<footer>
		<p class="mt-40 text-center font-bold"></p>
	</footer>
</div>

