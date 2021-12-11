package auctiongrpc

import (
	"context"
	"auction_api/auction/delivery/grpc/pb"

	gt "github.com/go-kit/kit/transport/grpc"
)


type GRPCServer struct{
	newAuction gt.Handler
	auctionData gt.Handler
}

func (s *GRPCServer) NewAuction(ctx context.Context, req *pb.NewAuctionRequest) (*pb.NewAuctionResponse, error){
	_, resp, err := s.newAuction.ServeGRPC(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp.(*pb.NewAuctionResponse), nil
}

func (s *GRPCServer) AuctionData(ctx context.Context, req *pb.AuctionDataRequest) (*pb.AuctionDataResponse, error){
	_, resp, err := s.auctionData.ServeGRPC(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp.(*pb.AuctionDataResponse), nil
}


func NewGRPCServer(endpoints Endpoints) pb.AuctionServiceServer {
	return &GRPCServer{
		newAuction: gt.NewServer(
			endpoints.NewAuction,
			decodeNewAuctionRequest,
			encodeNewAuctionResponse,
		),
		auctionData: gt.NewServer(
			endpoints.AuctionData,
			decodeAuctionDataRequest,
			encodeAuctionDataResponse,
		),
	}
}

func decodeNewAuctionRequest(_ context.Context, request interface{}) (interface{}, error){
	req := request.(*pb.NewAuctionRequest)
	return NewAuctionRequest{UserId: req.UserId, Title: req.Title, Type: req.Type, Status: req.Status, Date: req.Date}, nil
}

func encodeNewAuctionResponse(_ context.Context, response interface{}) (interface{}, error){
	resp := response.(NewAuctionResponse)
	return &pb.NewAuctionResponse{AuctionId: resp.ID}, nil
}

func decodeAuctionDataRequest(_ context.Context, request interface{}) (interface{}, error){
	req := request.(*pb.AuctionDataRequest)
	return AuctionDataRequest{ID: req.AuctionId}, nil
}

func encodeAuctionDataResponse(_ context.Context, response interface{}) (interface{}, error){
	resp := response.(AuctionDataResponse)
	return &pb.AuctionDataResponse{AuctionId: resp.ID, Title: resp.Title, Type: resp.Type, Status: resp.Status, Date: resp.Date}, nil
}