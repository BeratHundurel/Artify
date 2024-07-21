<script lang="ts">
	import type { Artwork } from '../types/artwork';

	export let data: {
		artworks: Artwork[];
	};

	let artworks: Artwork[] = data.artworks;
</script>

{#if artworks.length === 0}
	<p>No artworks found</p>
{:else}
	<div class="mx-auto max-w-xs sm:max-w-lg md:max-w-2xl lg:max-w-4xl xl:max-w-5xl">
		<div class="flex flex-col py-3">
			<h1
				class="border-b border-gray-400 pb-3 text-2xl font-semibold uppercase tracking-wide text-gray-800"
			>
				The Collection
			</h1>
			<p class="border-b border-gray-400 pb-8 pt-4 text-lg leading-8 text-gray-500">
				Explore thousands of artworks in the museum’s collection from our renowned icons to
				lesser-known works from every corner of the globe as well as our books, writings, reference
				materials, and other resources.
			</p>
			<div class="flex items-center gap-3 border-b border-red-900 py-2">
				<h2 class="py-2 text-xl font-medium uppercase tracking-wide text-gray-800">Artworks</h2>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
					fill="currentColor"
					class="size-5 text-gray-800"
				>
					<path
						fill-rule="evenodd"
						d="M16.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L14.69 12 7.72 5.03a.75.75 0 0 1 1.06-1.06l7.5 7.5Z"
						clip-rule="evenodd"
					/>
				</svg>
			</div>
			<div
				class="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
			>
				{#each artworks as artwork}
					<div class="flex flex-col items-center justify-center rounded-lg bg-gray-100 p-4">
						<h3 class="mt-2 text-lg font-semibold text-gray-800">{artwork.title}</h3>
						<p class="mt-1 text-sm text-gray-500">{artwork.artist_display}</p>
					</div>
				{/each}
			</div>
		</div>
	</div>
{/if}
