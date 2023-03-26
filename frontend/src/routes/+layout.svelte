<script>
	import TheNav from '../components/TheNav.svelte';
	import './app.css';
</script>

<div class="app mx-4 md:mx-8">
	<TheNav />

	<main>
		<slot />
	</main>

	<footer>
		<p class="mt-40 text-center font-bold">footer todo</p>
	</footer>
</div>

