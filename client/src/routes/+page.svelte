<script lang="ts">
	import * as Pagination from '$lib/components/ui/pagination';
	import { afterUpdate, onMount } from 'svelte';
	import { gsap } from 'gsap';
	export let data;
	afterUpdate(() => {
		const masonryItems = document.querySelectorAll('.masonry-item');
		gsap.from(masonryItems, {
			opacity: 0,
			y: 50,
			stagger: 0.1,
			duration: 1,
			ease: 'back.inOut'
		});
	});
</script>

{#if data.artworks.length === 0}
	<p>No artworks found</p>
{:else}
	<div
		class="mx-auto max-w-xs py-6 sm:max-w-lg md:max-w-2xl lg:max-w-4xl xl:max-w-5xl 2xl:max-w-7xl"
	>
		<div class="flex flex-col py-3">
			<h1
				class="border-b border-red-950 py-4 text-center text-4xl font-semibold uppercase tracking-wider text-red-950"
			>
				The Wall of Art
			</h1>
			<div class="masonry">
				{#each data.artworks as artwork}
					<div class="masonry-item flex flex-col items-start">
						<img src={artwork.full_image_url} alt={artwork.title} class="rounded-sm object-cover" />
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
