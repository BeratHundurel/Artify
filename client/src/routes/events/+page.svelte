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
		<div class="flex-lg-row flex flex-col">
			<div class="w--lg-1/5 w-full px-3 mb-3">
				{formattedDate}
			</div>
			<div class="w-lg-4/5 w-full px-3">
				{#each data.events as event}
					<div class="flex-lg-row mb-16 flex flex-col gap-6">
						<img src={event.image_url} alt={event.title} class="w-full w-lg-1/3 rounded-sm object-cover" />
						<div class="flex flex-col gap-4">
							<h2 class="text-foreground text-lg-xl text-lg font-medium">
								{event.title}
							</h2>
							<div class="text-muted-foreground flex flex-col gap-2 text-lg-lg text-base">
								{@html event.description}
							</div>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</Wrapper>
{/if}
