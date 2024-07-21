import type { PageServerLoad } from './$types';
import type { Artwork } from '../types/artwork'; // Adjust the import path as needed

export const load: PageServerLoad = async ({ fetch }) => {
	try {
		const res = await fetch('/api/artworks?page=1&limit=12');
		if (!res.ok) {
			throw new Error('Failed to fetch artworks');
		}
		const artworks: Artwork[] = await res.json();
		return {
			artworks
		};
	} catch (error) {
		console.error('Error loading artworks:', error);
		return {
			artworks: []
		};
	}
};
