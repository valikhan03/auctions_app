package auctionUsecase

import (
	"auctionservice/auction"
	"auctionservice/models"
)

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

func (a *AuctionUseCase) GetAuction(user_id, auction_id string) (*models.Auction, []string, error) {
	auction_data, err := a.repository.GetAuctionData(auction_id)
	if err != nil {
		return nil, nil, err
	}

	if auction_data.Owner == user_id {
		var participants []string
		participants, err = a.repository.GetAuctionParticipants(auction_id)
		if err != nil {
			return &auction_data, nil, nil
		}
		return &auction_data, participants, nil
	}

	return &auction_data, nil, nil
}
