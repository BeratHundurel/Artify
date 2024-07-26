import type { Event } from "../../types/event";

export async function fetchEvents(): Promise<Event[]> {
	const res = await fetch('http://localhost:5173/api/events');
	if (!res.ok) {
		throw new Error('Failed to fetch events from the server (HTTP ' + res.status + ')');
	}
	return res.json();
}
