<script lang="ts">
	import clsx from 'clsx';
	export let tag = 'section';
</script>

<svelte:element
	this={tag}
	{...$$restProps}
	class={clsx(
		`mx-auto max-w-xs px-3 py-3 sm:max-w-lg md:max-w-2xl lg:max-w-4xl lg:px-6 lg:py-6 xl:max-w-5xl 2xl:max-w-7xl`,
		$$props.class
	)}
>
	<slot />
</svelte:element>
