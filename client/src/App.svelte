<script>
	import Auth from './auth/auth.svelte'
	import Workspace from './workspace.svelte'
	import Footer from './footer.svelte'
	import { isLoggedStore } from './storeApp'
	let isLogged
	isLoggedStore.subscribe(value => { isLogged = value })
</script>
	
<div class="app">
	{#if !isLogged}
		<Auth />
	{:else}
		<Workspace />
	{/if}
	<Footer />
</div>

<style>
	.app {
		width: 100%;
		height: 100%;
	}
	/*global style for body*/
	:global(body) {
		margin: 0;
		padding: 0;
		width: 100%;
		height: 100%;
	}
</style>