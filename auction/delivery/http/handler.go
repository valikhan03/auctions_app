package auctionhttp

import (
	"auctionservice/auction"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	UseCase auction.UseCase
}

func NewHandler(usecase auction.UseCase) *Handler{
	return &Handler{
		UseCase: usecase,
	}
}


type NewAuctionInput struct{
	title string
	owner_id string
}


func (h *Handler) NewAuction(c *gin.Context){
	
	var title string
	c.BindJSON(&title)
	

	
	//h.UseCase.CreateAuction()*/
}