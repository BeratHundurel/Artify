<script lang="ts">
	import Wrapper from '$lib/components/ui/Wrapper.svelte';
	import { onMount } from 'svelte';

	let paragraph: HTMLParagraphElement;
	let image: HTMLImageElement;
	let vector: HTMLDivElement;
	let vector2: HTMLDivElement;
	let vector3: HTMLDivElement;
	let mobileVector: HTMLDivElement;
	onMount(() => {
		if (paragraph) {
			const paragraphHeight = paragraph.offsetHeight;
			vector.style.bottom = `${paragraphHeight + 4}px`;
			vector.style.height = `calc(100% - ${paragraphHeight}px)`;
			vector3.style.bottom = `${paragraphHeight - 4}px`;
			vector3.style.height = `calc(100% - ${paragraphHeight}px + 28px)`;
			if (mobileVector) {
				mobileVector.style.bottom = `${paragraphHeight + 4}px`;
			}
		}
		if (image) {
			const imageWidth = image.offsetWidth;
			vector2.style.width = `${imageWidth}px`;
		}
	});
</script>

<div class="flex flex-col">
	<section
		class="relative flex h-[80svh] w-full flex-col items-center justify-center"
		aria-label="Banner"
		aria-describedby="Banner of the home page"
	>
		<img src="/artbanner.png" alt="Banner" class="relative h-full w-full" />
		<h1
			class="absolute bottom-[5%] z-20 w-3/4 text-center text-3xl uppercase tracking-wider text-slate-100 lg:w-1/2 lg:text-5xl"
		>
			We preserve History To build Future
		</h1>
		<div
			class="absolute bottom-0 z-10 flex h-1/3 w-full items-center justify-center bg-black opacity-70"
		></div>
	</section>
	<Wrapper tag="section" class="relative grid grid-cols-1 gap-8 lg:grid-cols-3">
		<div class="flex flex-col gap-6">
			<h1 class="relative z-10 text-4xl !leading-[1.2] tracking-wider text-slate-100 lg:text-8xl">
				The World’s Leading Museum
				<div
					class="absolute bottom-0 h-[1px] w-full bg-gray-300 lg:bottom-1/4"
					bind:this={vector2}
				></div>
			</h1>
			<p class="text-lg leading-normal text-slate-100 lg:text-2xl">
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Scelerisque et, condimentum iaculis
				ac. Enim leo sit vestibulum vulputate placerat vulputate. Quis egestas pellentesque faucibus
				scelerisque consectetur amet. Imperdiet in posuere tempus enim sit semper turpis faucibus
				lectus. Eget et tempor dolor, amet,
			</p>
		</div>
		<div class="relative z-10 flex flex-col gap-6 lg:col-span-2">
			<img class="object cover z-10 w-full" alt="about" src="/about.png" bind:this={image} />
			<p
				class="z-10 text-lg leading-normal text-slate-100 lg:ps-4 lg:text-2xl"
				bind:this={paragraph}
			>
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Scelerisque et, condimentum iaculis
				ac. Enim leo sit vestibulum vulputate placerat vulputate. Quis egestas pellentesque faucibus
				scelerisque consectetur amet. Imperdiet in posuere tempus enim sit semper turpis faucibus
				lectus. Eget et tempor dolor, amet, vel sit tellus sem. Lectus sed leo neque ut id nunc,
				amet. Ante non, tortor nisl, vitae orci lacus ipsum vehicula eu. Bibendum.
			</p>
			<div class="absolute left-12 w-[1px] bg-gray-300" bind:this={vector}></div>
			<div class="absolute right-12 w-[1px] bg-gray-300" bind:this={vector3}></div>
			<div
				class="absolute left-[calc(20%-24px)] top-4 h-[1px] w-[calc(75%+4rem)] bg-gray-300 lg:hidden"
			></div>
			<div
				class="absolute right-[calc(20%-24px)] h-[1px] w-[calc(75%+4rem)] bg-gray-300 lg:hidden"
				bind:this={mobileVector}
			></div>
		</div>
		<div
			class="absolute left-[calc(33.33%-24px)] top-16 hidden h-[1px] w-[calc(66%+4rem)] bg-gray-300 lg:block"
		></div>
	</Wrapper>
</div>
