package types

type Event struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	TitleDisplay     string `json:"title_display"`
	ImageURL         string `json:"image_url"`
	HeroCaption      string `json:"hero_caption"`
	ShortDescription string `json:"short_description"`
	ListDescription  string `json:"list_description"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	DateDisplay      string `json:"date_display"`
}
