import type { PageServerLoad } from './$types';
import type { Artwork } from '../types/artwork'; // Adjust the import path as needed

export const load: PageServerLoad = async ({ url, fetch }) => {
	try {
		const page = url.searchParams.get('page') || '1';
		const res = await fetch(`/api/artworks?page=${page}`);
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
