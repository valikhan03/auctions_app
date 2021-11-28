package auctionhttp

import (
	"auctionservice/auction"

	"github.com/gin-gonic/gin"
)

func RegisterAuctionHttpEndpoints(router *gin.RouterGroup, uc auction.UseCase) {
	h := NewHandler(uc)

	auctionEndpoints := router.Group("/auctions")
	{
		auctionEndpoints.POST("/new_auction", h.NewAuction)
	}

}
