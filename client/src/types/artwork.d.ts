export interface Artwork {
	classification_titles: string[];
	short_description: string;
	is_public_domain: boolean;
	thumbnail: Thumbnail;
	artist_display: string;
	artist_title: string;
	title: string;
	artist_id: number;
	artwork_type_title: string;
	date_display: string;
	id: number;
	image_id: string;
	artwork_type_id: number;
}

export interface Thumbnail {
	alt_text: string;
	width: number;
	lqip: string;
	height: number;
}
