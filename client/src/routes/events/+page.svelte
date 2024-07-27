<script lang="ts">
	import Wrapper from '$lib/components/ui/Wrapper.svelte';

	export let data;
</script>

{#if data.events.length === 0}
	<p>No events found</p>
{:else}
	<Wrapper>
		<div class="flex flex-row">
			<div class="w-1/5 px-3">
				{new Date().getDay()}
				{new Date().getMonth()}
			</div>
			<div class="w-4/5 px-3">
				{#each data.events as event}
					<div class="mb-16 flex gap-6">
						<img src={event.image_url} alt={event.title} class="w-1/3 rounded-sm object-cover" />
						<div class="flex flex-col gap-4">
							<h2 class="text-foreground text-xl font-medium">
								{event.title}
							</h2>
							<div class="text-muted-foreground flex flex-col gap-2 text-lg">
								{@html event.description}
							</div>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</Wrapper>
{/if}
