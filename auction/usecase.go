package auction

import "auctionservice/models"

type UseCase interface {
	CreateAuction(user_id, auctionTitle string) (string, error)
	InviteParticipant(user_id, auction_id string) error
	EnrollToAuction()
	GetAuction(user_id, auction_id string) (*models.Auction, []string, error)
}
