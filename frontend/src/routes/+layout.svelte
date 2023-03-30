<script>
	import TheNav from '../components/TheNav.svelte';
	import './app.css';
</script>

<div class="">
	<TheNav />

	<main class="mx-auto max-w-5xl px-4 lg:px-8">
		<slot />
	</main>

	<footer>
		<p class="mt-40 text-center font-bold">footer todo</p>
	</footer>
</div>

