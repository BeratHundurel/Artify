<script lang="ts">
	import Wrapper from './Wrapper.svelte';
</script>

<Wrapper tag="Header" class="h-[20svh]">
	<div class="hidden items-center justify-center gap-20 text-center lg:flex">
		<a href="/" class="w-24 text-lg text-slate-100 first-letter:uppercase"> Home </a>
		<a href="/gallery" class="w-24 text-lg text-slate-100 first-letter:uppercase"> Gallery </a>
		<div class="flex flex-col items-center justify-center">
			<img src="/logo.png" alt="Logo" class="object-fit-cover" />
			<h1 class="text-xl">Museum</h1>
		</div>
		<a href="/events" class="w-24 text-lg text-slate-100 first-letter:uppercase"> Events </a>
		<a href="/history" class="w-24 text-lg text-slate-100 first-letter:uppercase"> History </a>
	</div>
	<div class="flex items-center justify-between lg:hidden">
		<div>
			<img src="/logo.png" alt="Logo" class="object-fit-cover" />
			<h3 class="text-lg">Museum</h3>
		</div>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="lucide lucide-align-justify"
			><line x1="3" x2="21" y1="6" y2="6" /><line x1="3" x2="21" y1="12" y2="12" /><line
				x1="3"
				x2="21"
				y1="18"
				y2="18"
			/></svg
		>
	</div>
</Wrapper>
