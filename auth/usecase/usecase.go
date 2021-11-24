package authUsecase

import (
	"auctionservice/auth"
	"auctionservice/models"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

type AuthUseCase struct {
	userRepos      auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	user_repos auth.UserRepository,
	hash_salt string,
	signing_key []byte,
	token_tls_seconds time.Duration) *AuthUseCase {

	return &AuthUseCase{
		userRepos:      user_repos,
		hashSalt:       hash_salt,
		signingKey:     signing_key,
		expireDuration: time.Second * token_tls_seconds,
	}
}

func (a *AuthUseCase) SignUp(email, firstname, lastname, password string) error {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepos.CreateUser(user)
}

func (a *AuthUseCase) ParseToken(accessToken string) (*models.User, error) {
	return nil, errors.New("")
}

func (a *AuthUseCase) SignIn(email, password string) (string, error) {
	return "", nil
}
