package models

type Product struct {
	Name   string   `json:"name"`
	Images []string `json:"images"`
}

type Lot struct {
	LotID       string    `json:"lot_id"`
	Owner       string    `json:"owner_id"`
	Auction     string    `json:"auction_id"`
	Products    []Product `json:"product"`
	Start_price int       `json:"start_price"`
}
