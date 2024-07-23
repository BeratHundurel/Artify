<script lang="ts">
	import * as Pagination from '$lib/components/ui/pagination';
	import { afterUpdate, onMount } from 'svelte';
	import { gsap } from 'gsap';
	import Wrapper from '$lib/components/ui/Wrapper.svelte';
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
	<Wrapper tag="section">
		<div class="flex flex-col py-3">
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
					<a href="?page={currentPage}">
						<Pagination.NextButton />
					</a>
				</Pagination.Item>
			</Pagination.Content>
		</Pagination.Root>
	</Wrapper>
{/if}
