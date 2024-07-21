package types

type ArtworkResponse struct {
	ClassificationTitles []string  `json:"classification_titles"`
	ShortDescription     string    `json:"short_description"`
	IsPublicDomain       bool      `json:"is_public_domain"`
	Thumbnail            Thumbnail `json:"thumbnail"`
	ArtistDisplay        string    `json:"artist_display"`
	ArtistTitle          string    `json:"artist_title"`
	Title                string    `json:"title"`
	ArtistID             int       `json:"artist_id"`
	ArtworkTypeTitle     string    `json:"artwork_type_title"`
	DateDisplay          string    `json:"date_display"`
	ID                   int       `json:"id"`
	ImageID              string    `json:"image_id"`
	ArtworkTypeID        int       `json:"artwork_type_id"`
}

type Artwork struct {
	ClassificationTitles []string  `json:"classification_titles"`
	ShortDescription     string    `json:"short_description"`
	IsPublicDomain       bool      `json:"is_public_domain"`
	Thumbnail            Thumbnail `json:"thumbnail"`
	ArtistDisplay        string    `json:"artist_display"`
	ArtistTitle          string    `json:"artist_title"`
	Title                string    `json:"title"`
	ArtistID             int       `json:"artist_id"`
	ArtworkTypeTitle     string    `json:"artwork_type_title"`
	DateDisplay          string    `json:"date_display"`
	ID                   int       `json:"id"`
	ImageID              string    `json:"image_id"`
	ArtworkTypeID        int       `json:"artwork_type_id"`
	FullImageURL         string    `json:"full_image_url"`
}
