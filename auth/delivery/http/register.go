package deliveryhttp

import (
	"auctionservice/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHadler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
	}
}
