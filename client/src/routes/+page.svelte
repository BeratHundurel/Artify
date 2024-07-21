<script lang="ts">
	import * as Pagination from '$lib/components/ui/pagination';
	export let data;
</script>

{#if data.artworks.length === 0}
	<p>No artworks found</p>
{:else}
	<div
		class="mx-auto max-w-xs py-6 sm:max-w-lg md:max-w-2xl lg:max-w-4xl xl:max-w-5xl 2xl:max-w-7xl"
	>
		<div class="flex flex-col py-3">
			<h1
				class="border-b border-gray-400 pb-3 text-2xl font-semibold uppercase tracking-wide text-gray-800"
			>
				The Collection
			</h1>
			<p class="pb-8 pt-4 text-lg leading-8 text-gray-500">
				Explore thousands of artworks in the museum’s collection from our renowned icons to
				lesser-known works from every corner of the globe as well as our books, writings, reference
				materials, and other resources.
			</p>
			<div class="mt-4 h-[1px] w-full bg-gray-300 lg:mt-8"></div>
			<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
				{#each data.artworks as artwork}
					<div class="grid-item flex flex-col items-start justify-center pl-4 pr-4 pt-4">
						<img src={artwork.full_image_url} alt={artwork.title} class="rounded-sm object-cover" />
						<h3 class="mt-4 text-base font-medium text-gray-700">{artwork.title}</h3>
						<p class="mt-2 text-sm text-gray-600">
							{artwork.artist_display}
						</p>
						<div class="mt-10 h-[1px] w-full bg-gray-300" />
					</div>
				{/each}
			</div>
		</div>
		<Pagination.Root count={120} perPage={12} let:pages let:currentPage>
			<Pagination.Content>
				<Pagination.Item>
					<Pagination.PrevButton />
				</Pagination.Item>
				{#each pages as page (page.key)}
					{#if page.type === 'ellipsis'}
						<Pagination.Item>
							<Pagination.Ellipsis />
						</Pagination.Item>
					{:else}
						<a href="?page={page.value}">
							<Pagination.Item>
								<Pagination.Link {page} isActive={currentPage == page.value}>
									{page.value}
								</Pagination.Link>
							</Pagination.Item>
						</a>
					{/if}
				{/each}
				<Pagination.Item>
					<Pagination.NextButton />
				</Pagination.Item>
			</Pagination.Content>
		</Pagination.Root>
	</div>
{/if}

<style>
	.grid-item:not(:nth-child(4n)):not(:last-child) {
		border-right: 1px solid #e5e7eb; /* Tailwind's border-gray-200 */
	}
</style>
