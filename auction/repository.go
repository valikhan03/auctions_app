package auction

import "auctionservice/models"

type AuctionRepository interface {
	NewAuction(user_id, auction_title string) (string, error)
	AddParticipant(auction_id string, user_id string) error
	GetAuctionData(auction_id string) (models.Auction, error)
	GetAuctionParticipants(auction_id string) ([]string, error)
}
