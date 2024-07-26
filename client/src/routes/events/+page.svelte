<script lang="ts">
	export let data;
</script>

{#if data.events.length === 0}
	<p>No events found</p>
{:else}
	<div class="flex flex-row">
		<div class="w-1/5 px-3">
			{new Date().getDay()}
			{new Date().getMonth()}
		</div>
		<div class="w-4/5 px-3">
			{#each data.events as event}
				<div class="flex gap-4 mb-8">
					<img src={event.image_url} alt={event.title} class="w-1/4 object-cover" />
					<div class="flex flex-col gap-4">
						<h2 class="text-foreground text-xl font-medium">{event.title}</h2>
						<p class="text-muted-foreground text-base font-normal">{event.short_description}</p>
						{@html event.description}
					</div>
				</div>
			{/each}
		</div>
	</div>
{/if}
