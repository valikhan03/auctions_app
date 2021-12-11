package auctionUsecase

import (
	"auction_api/auction"
	"auction_api/models"
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

func (a *AuctionUseCase) CreateAuction(user_id, auctionTitle, auctionType, status, date string) (string, error) {
	auction_id, err := a.repository.NewAuction(auctionTitle, auctionType, status, date)
	return auction_id, err
}

func (a AuctionUseCase) InviteParticipant(user_id, auction_id string) error {
	return nil
}

func (a *AuctionUseCase) EnrollToAuction() {

}

func (a *AuctionUseCase) GetAuction(auction_id string) (*models.Auction, error) {
	auction_data, err := a.repository.GetAuctionData(auction_id)
	if err != nil {
		return nil, err
	}

	return &auction_data, nil
}



func (a *AuctionUseCase) GetAllPublicAuctions() (*[]models.Auction, error){
	auctions, err := a.repository.GetAllPublicAuctions()
	if err != nil{
		return nil, err
	}
	return auctions, nil
}