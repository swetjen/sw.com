<script>
	import TheNav from '../components/TheNav.svelte';
	import ContactModal from '../components/ContactModal.svelte';
	import './app.css';
</script>

<div class="max-w-5xl mx-auto px-4 md:px-8">
	<TheNav />
	<ContactModal />

	<main>
		<slot />
	</main>

	<footer>
		<p class="mt-40 text-center font-bold">footer todo</p>
	</footer>
</div>

