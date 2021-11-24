package auction


type UseCase interface{
	CreateAuction(user_id, auctionTitle string) (string, error)
	InviteParticipant(user_id, auction_id string) error
	EnrollToAuction()
	GetAuction()
}

