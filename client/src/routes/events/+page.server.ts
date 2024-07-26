import { fetchEvents } from '$lib/server/event';
import type { Event } from '../../types/event';
import type { PageServerLoad } from '../gallery/$types';

export const load: PageServerLoad = async () => {
	try {
		const events: Event[] = await fetchEvents();
		return {
			events
		};
	} catch (error) {
		console.error('Error loading artworks:', error);
		return {
			events: []
		};
	}
};
