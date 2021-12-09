package auth

import(
	"auctionservice/models"
)


type UserRepository interface{
	CreateUser(user *models.User) error
	GetUserID(email, password string) (string, error)
}