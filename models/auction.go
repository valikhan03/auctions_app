package models

type Auction struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`   //private public
	Status string `json:"status"` //started/ended/not started
	Date   string `json:"Date"`
	Owner  string `json:"owner_id"`
}

type AuctionParticipants struct {
	Id           string   `json:"id"`
	Participants []string `json:"participants"`
}
