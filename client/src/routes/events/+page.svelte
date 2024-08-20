<script lang="ts">
	import Wrapper from '$lib/components/ui/Wrapper.svelte';
	const formattedDate = new Intl.DateTimeFormat('tr-TR', {
		day: '2-digit',
		month: '2-digit',
		year: 'numeric'
	}).format(new Date());
	export let data;
</script>

{#if data.events.length === 0}
	<p>No events found</p>
{:else}
	<Wrapper>
		<div class="lg:flex-row flex flex-col">
			<div class="mb-3 w-full px-3 lg:w-1/5">
				{formattedDate}
			</div>
			<div class="w-full px-3 lg:w-4/5">
				{#each data.events as event}
					<div class="mb-16 flex flex-col gap-6 lg:flex-row">
						<img
							src={event.image_url}
							alt={event.title}
							class="lg:w-1/3 w-full rounded-sm object-cover"
						/>
						<div class="flex flex-col gap-4">
							<h2 class="text-foreground text-lg font-medium lg:text-xl">
								{event.title}
							</h2>
							<div class="text-muted-foreground flex flex-col gap-2 text-base lg:text-lg">
								{@html event.description}
							</div>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</Wrapper>
{/if}
