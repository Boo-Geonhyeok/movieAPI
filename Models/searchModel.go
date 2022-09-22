package models

type TitleResult struct {
	Meta struct {
		Operation     string  `json:"operation"`
		RequestID     string  `json:"requestId"`
		ServiceTimeMs float64 `json:"serviceTimeMs"`
	} `json:"@meta"`
	Type          string `json:"@type"`
	PaginationKey string `json:"paginationKey"`
	Results       []struct {
		Disambiguation string `json:"disambiguation,omitempty"`
		ID             string `json:"id"`
		Image          struct {
			Height int    `json:"height"`
			ID     string `json:"id"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image,omitempty"`
		Title     string `json:"title"`
		TitleType string `json:"titleType"`
		Year      int    `json:"year"`
	} `json:"results"`
	TotalMatches int `json:"totalMatches"`
}
