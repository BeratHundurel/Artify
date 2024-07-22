import type { Artwork } from '../../types/artwork';

export async function fetchArtworks(page: string): Promise<Artwork[]> {
	const res = await fetch('http://localhost:5173/api/artworks?page=' + page);
	if (!res.ok) {
		throw new Error('Failed to fetch artworks from the server (HTTP ' + res.status + ')');
	}
	return res.json();
}
