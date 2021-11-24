package auctionUsecase

import "auctionservice/auction"

type AuctionUseCase struct {
	repository auction.AuctionRepository

	//other configs
}

func NewAuctionUseCase(repos auction.AuctionRepository) *AuctionUseCase {
	return &AuctionUseCase{
		repository: repos,
	}
}

func (a *AuctionUseCase) CreateAuction(user_id, auctionTitle string) (string, error) {
	auction_id, err := a.repository.NewAuction(user_id, auctionTitle)
	return auction_id, err
}

func (a AuctionUseCase) InviteParticipant(user_id, auction_id string) error {
	return nil
}

func (a *AuctionUseCase) EnrollToAuction() {

}

func (a *AuctionUseCase) GetAuction() {}
