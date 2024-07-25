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
	baseURL       = "https://api.artic.edu/api/v1/artworks/search"
	artworkFields = "id,title,thumbnail,date_display,artist_display,place_of_origin_description,short_description,artwork_type_title,artwork_type_id,artist_id,artist_title,classification_titles,image_id,is_public_domain"
)

// Initialize an HTTP client with a timeout
var client = &http.Client{Timeout: 10 * time.Second}

type Artworks struct {
	Data []types.Artwork `json:"data"`
}

func buildURL(page, limit int, endpoint, fields string) (string, error) {
	u, err := url.Parse(endpoint)
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
	url, err := buildURL(page, limit, baseURL, artworkFields)
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

	var artworks Artworks
	err = json.Unmarshal(body, &artworks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return artworks.Data, nil
}
