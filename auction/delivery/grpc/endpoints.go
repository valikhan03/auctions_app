package auctiongrpc

import (
	"auction_api/auction"
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	NewAuction  endpoint.Endpoint
	AuctionData endpoint.Endpoint
}

type NewAuctionRequest struct {
	UserId string
	Title  string `json:"title"`
	Type   string `json:"type"`   //private public
	Status string `json:"status"` //started/ended/not started
	Date   string `json:"Date"`
}

type NewAuctionResponse struct {
	ID string
}

type AuctionDataRequest struct {
	ID string
}

type AuctionDataResponse struct {
	ID     string
	Title  string `json:"title"`
	Type   string `json:"type"`   //private public
	Status string `json:"status"` //started/ended/not started
	Date   string `json:"Date"`
}


func MakeEndpoints(uc auction.UseCase) Endpoints{
	return Endpoints{
	}
}

func makeNewAuctionEnpoints(uc auction.UseCase) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewAuctionRequest)

		result, err := uc.CreateAuction(req.UserId, req.Title, req.Type, req.Status, req.Date)
		if err != nil{
			log.Println(err)
			return nil, err
		}

		return NewAuctionResponse{ID: result}, nil
	}
}