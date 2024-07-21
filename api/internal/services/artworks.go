package services

import (
	"api/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.artic.edu/api/v1/artworks/search"
	fields  = "id,title,thumbnail,date_display,artist_display,place_of_origin_description,short_description,artwork_type_title,artwork_type_id,artist_id,artist_title,classification_titles,image_id,is_public_domain"
)

// Initialize an HTTP client with a timeout
var client = &http.Client{Timeout: 10 * time.Second}

type ArtworksResponse struct {
	Data []types.ArtworkResponse `json:"data"`
}

func buildURL(page, limit int) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("query[term][is_public_domain]", "true")
	params.Add("fields", fields)
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("limit", fmt.Sprintf("%d", limit))
	u.RawQuery = params.Encode()

	return u.String(), nil
}

func FetchArtworks(page, limit int) ([]types.Artwork, error) {
	url, err := buildURL(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var artworksResponse ArtworksResponse
	err = json.Unmarshal(body, &artworksResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	artworks := make([]types.Artwork, 0, len(artworksResponse.Data))
	for _, artworkResponse := range artworksResponse.Data {
		artwork := types.Artwork{
			ClassificationTitles: artworkResponse.ClassificationTitles,
			ShortDescription:     artworkResponse.ShortDescription,
			IsPublicDomain:       artworkResponse.IsPublicDomain,
			Thumbnail:            artworkResponse.Thumbnail,
			ArtistDisplay:        artworkResponse.ArtistDisplay,
			ArtistTitle:          artworkResponse.ArtistTitle,
			Title:                artworkResponse.Title,
			ArtistID:             artworkResponse.ArtistID,
			ArtworkTypeTitle:     artworkResponse.ArtworkTypeTitle,
			DateDisplay:          artworkResponse.DateDisplay,
			ID:                   artworkResponse.ID,
			ImageID:              artworkResponse.ImageID,
			ArtworkTypeID:        artworkResponse.ArtworkTypeID,
			FullImageURL:         fmt.Sprintf("https://www.artic.edu/iiif/2/%s/full/843,/0/default.jpg", artworkResponse.ImageID),
		}
		artworks = append(artworks, artwork)
	}

	return artworks, nil
}
