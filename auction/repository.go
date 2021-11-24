package auction

type AuctionRepository interface {
	NewAuction(user_id, auction_title string) (string, error)
	AddParticipant(auction_id string, user_id string) error
	GetAuction()
}
