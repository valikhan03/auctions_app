package auction

import "auction_api/models"

type AuctionRepository interface {
	NewAuction(auctionTitle, auctionType, auctionStatus, auctionDate string) (string, error)
	AddParticipant(auction_id string, user_id string) error
	GetAuctionData(auction_id string) (models.Auction, error)
	GetAuctionParticipants(auction_id string) ([]string, error)
	GetAllPublicAuctions() (*[]models.Auction, error)
}
