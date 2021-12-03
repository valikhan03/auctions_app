package auctionhttp

import (
	"auctionservice/auction"
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
