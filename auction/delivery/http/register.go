package auctionhttp

import (
	"auctionservice/auction"

	"github.com/gin-gonic/gin"
)

func RegisterAuctionHttpEndpoints(router *gin.Engine, uc auction.UseCase) {
	h := NewHandler(uc)

	
		router.POST("/new_auction", h.NewAuction)
		router.GET("/:id", h.GetAuctionData)

}
