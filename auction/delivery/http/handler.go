package auctionhttp

import (
	"auctionservice/auction"
	"auctionservice/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase auction.UseCase
}

func NewHandler(usecase auction.UseCase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}

func (h *Handler) NewAuction(c *gin.Context) {
	var title string
	c.BindJSON(&title)

	user_id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	auction_id, err := h.UseCase.CreateAuction(user_id.(string), title)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, auction_id)
}

type AuctionData struct {
	Auction      models.Auction `json:"auction"`
	Participants []string       `json:"participants"`
}

func (h *Handler) GetAuctionData(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if exists == false{
		c.AbortWithStatus(http.StatusBadRequest)
	}
	auction_id, ok := c.Params.Get("id")
	if ok == false {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	auction, participants, err := h.UseCase.GetAuction(user_id.(string), auction_id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var auctionData  = &AuctionData{
		Auction: *auction,
		Participants: participants,
	}  
	c.JSON(http.StatusOK ,auctionData)

}
