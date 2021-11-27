package models

type Auction struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner_id"`
}

type AuctionParticipants struct {
	Id           string   `json:"id"`
	Participants []string `json:"participants"`
}
