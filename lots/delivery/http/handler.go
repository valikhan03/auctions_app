package lotshttp

import (
	"auction_api/lots"
	"auction_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LotsHadler struct{
	UseCase lots.UseCase
}


func NewLotsHandler(usecase lots.UseCase) *LotsHadler{
	return &LotsHadler{
		UseCase: usecase,
	}
}


func NewLot(c *gin.Context){
	var lot models.Lot

	err := c.BindJSON(&lot)
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	

}