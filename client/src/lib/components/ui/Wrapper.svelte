<script lang="ts">
	import clsx from 'clsx';
	export let tag = 'section';
</script>

<svelte:element
	this={tag}
	{...$$restProps}
	class={clsx(`px-3 py-3 lg:px-6 lg:py-6`, $$props.class)}
>
	<div
		class="mx-auto max-w-xs py-6 sm:max-w-lg md:max-w-2xl lg:max-w-4xl xl:max-w-5xl 2xl:max-w-7xl"
	>
		<slot />
	</div>
</svelte:element>
