package models

type Auction struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Owner        string   `json:"owner_id"`
	Participants []string `json:"participants"`
}
