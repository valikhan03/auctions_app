package auction

import "auction_api/models"

type UseCase interface {
	CreateAuction(user_id, auctionTitle, auctionType, status, date string) (string, error)
	InviteParticipant(user_id, auction_id string) error
	EnrollToAuction()
	GetAuction(auction_id string) (*models.Auction, error)
	GetAllPublicAuctions() (*[]models.Auction, error)
}
