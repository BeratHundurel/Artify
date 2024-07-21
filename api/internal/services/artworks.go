package services

import (
	"api/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.artic.edu/api/v1/artworks/search"
	fields  = "id,title,thumbnail,date_display,artist_display,place_of_origin_description,short_description,artwork_type_title,artwork_type_id,artist_id,artist_title,classification_titles,image_id,is_public_domain"
)

type ArtworksResponse struct {
	Data []types.Artworks `json:"data"`
}

func buildURL(page, limit int) string {
	return fmt.Sprintf("%s?query[term][is_public_domain]=true&fields=%s&page=%d&limit=%d", baseURL, fields, page, limit)
}

func FetchArtworks(page, limit int) ([]types.Artworks, error) {
	url := buildURL(page, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artworksResponse ArtworksResponse
	err = json.Unmarshal(body, &artworksResponse)
	if err != nil {
		return nil, err
	}

	return artworksResponse.Data, nil
}
