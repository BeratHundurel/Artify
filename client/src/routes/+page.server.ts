import type { PageServerLoad } from './$types';
import type { Artwork } from '../types/artwork'; // Adjust the import path as needed
import { fetchArtworks } from '$lib/server/artwork';

export const load: PageServerLoad = async ({ url }) => {
	try {
		const page = url.searchParams.get('page') || '1';
		const artworks: Artwork[] = await fetchArtworks(page);
		return {
			artworks,
		};
	} catch (error) {
		console.error('Error loading artworks:', error);
		return {
			artworks: []
		};
	}
};
