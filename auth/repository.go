package auth

import(
	"auctionservice/models"
)


type UserRepository interface{
	CreateUser(user *models.User) error
	GetUser(email, password string) (*models.User, error)
}