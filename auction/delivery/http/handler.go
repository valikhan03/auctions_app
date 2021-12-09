package auctionhttp

import (
	"auctionservice/auction"
	"auctionservice/models"
	"fmt"
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

type Title struct{
	Title string `json:"title"`
}

func (h *Handler) NewAuction(c *gin.Context) {
	var title Title
	c.BindJSON(&title)

	fmt.Println("handler-", title.Title)

	user_id, err := c.Cookie("userID")
	fmt.Println(user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	auction_id, err := h.UseCase.CreateAuction(user_id, title.Title)
	if err != nil {
		fmt.Println(err)
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
	user_id, err := c.Cookie("userID")
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	auction_id, ok := c.Params.Get("id")
	fmt.Println(auction_id)
	if ok == false {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	auction, participants, err := h.UseCase.GetAuction(user_id, auction_id)
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
