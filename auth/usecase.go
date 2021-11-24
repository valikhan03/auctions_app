package auth

import "auctionservice/models"

type UseCase interface {
	SignUp(email, password, firstname, lastname string) error
	SignIn(email, password string) (string, error)
	ParseToken(accessToken string) (*models.User, error)
}


