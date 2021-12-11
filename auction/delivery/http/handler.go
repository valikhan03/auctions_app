package auctionhttp

import (
	"auction_api/auction"
	"auction_api/models"
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

type AuctionData struct {
	Title  string `json:"title"`
	Type   string `json:"type"`   //private public
	Status string `json:"status"` //started/ended/not started
	Date   string `json:"Date"`
}

func (h *Handler) NewAuction(c *gin.Context) {
	var newAuction AuctionData
	c.BindJSON(&newAuction)

	user_id, err := c.Cookie("userID")
	fmt.Println(user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	auction_id, err := h.UseCase.CreateAuction(user_id, newAuction.Title, newAuction.Type, newAuction.Status, newAuction.Date)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, auction_id)
}

type AuctionFullData struct {
	Auction      models.Auction `json:"auction"`
	Participants []string       `json:"participants"`
}

func (h *Handler) GetAuctionData(c *gin.Context) {

	auction_id, ok := c.Params.Get("id")
	fmt.Println(auction_id)
	if ok == false {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	auction, err := h.UseCase.GetAuction(auction_id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var auctionFullData = &AuctionFullData{
		Auction:      *auction,
	}
	c.JSON(http.StatusOK, auctionFullData)

}

func (h *Handler) GetAllPublicAuctions(c *gin.Context) {
	auctions, err := h.UseCase.GetAllPublicAuctions()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, auctions)
}
