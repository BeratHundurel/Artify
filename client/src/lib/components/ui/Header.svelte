<script lang="ts">
	import Wrapper from './Wrapper.svelte';
</script>

<Wrapper tag="Header" class="h-[20vh]">
	<div class="flex items-center justify-center gap-20 text-center">
		<a href="/" class="w-24 text-slate-100 text-lg first-letter:uppercase"> Home </a>
		<a href="/" class="w-24 text-slate-100 text-lg first-letter:uppercase"> Gallery </a>
		<div class="flex flex-col items-center justify-center">
			<img src="/logo.png" alt="Logo" class="object-fit-cover" />
			<h1 class="text-xl">Museum</h1>
		</div>
		<a href="/" class="w-24 text-slate-100 text-lg first-letter:uppercase"> Events </a>
		<a href="/" class="w-24 text-slate-100 text-lg first-letter:uppercase"> History </a>
	</div>
</Wrapper>
