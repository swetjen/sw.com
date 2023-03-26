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
		<p class="my-5">todo</p>
	</footer>
</div>

