package services

import (
	"api/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	eventURL    = "https://api.artic.edu/api/v1/events"
	eventFields = "id,title,title_display,image_url,hero_caption,short_description,list_description,description,location,start_date,end_date,start_time,end_time,date_display"
)

type Events struct {
	Data []types.Event `json:"data"`
}

func FetchEvents(page, limit int) ([]types.Event, error) {
	url, err := buildURL(page, limit, eventURL, eventFields)
	if err != nil {
		return nil, err
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

	var events Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return events.Data, nil
}
