package lotsUsecase

import (
	"auctionservice/lots"
	"auctionservice/models"
)

type LotsUseCase struct {
	Repository lots.LotsRepository
}

func NewAuctionUseCase(repos lots.LotsRepository) *LotsUseCase{
	return &LotsUseCase{
		Repository: repos,
	}
}


func NewLot(lot models.Lot){

}

func AddProductToLot(product models.Product, lot_id string, user_id string){

}

func GetAllLots(auction_id string){

}

func GetLot(lot_id string){

}

func BuyLot(lot_id string){
	
}