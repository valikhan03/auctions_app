package deliveryhttp

import (
	"auctionservice/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase auth.UseCase
}

func NewHadler(usecase auth.UseCase) *Handler{
	return &Handler{
		UseCase: usecase,
	}
}

type SignUpInput struct{
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}


func (h *Handler) SignUp(c *gin.Context){
	input := new(SignUpInput)
	
	if err := c.BindJSON(&input); err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.UseCase.SignUp(input.Email, input.Password, input.Firstname, input.Lastname); err != nil{
		c.AbortWithStatus(http.StatusInternalServerError);
		return
	}

	c.Status(http.StatusOK)
}