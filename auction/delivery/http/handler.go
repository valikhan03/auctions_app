package deliveryhttp

import (
	"auctionservice/auction"
	//"log"
	//"net/http"

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

/*
type NewAuctionInput struct{
	title string
	owner_id string
}
*/

func (h *Handler) NewAuction(c *gin.Context){
	/*
	var title string
	c.BindJSON(&title)
	cookie, err := c.Request.Cookie("access-cookie")
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	cookie_content := cookie.Value

	
	//h.UseCase.CreateAuction()*/
}