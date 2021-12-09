package authUsecase

import (
	"auctionservice/auth"
	"auctionservice/models"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthUseCase struct {
	userRepos  auth.UserRepository
	hashSalt   string
	signingKey []byte
	expireTime time.Duration
}

func NewAuthUseCase(
	user_repos auth.UserRepository,
	hash_salt string,
	signing_key []byte,
	tokenTLSSeconds time.Duration) *AuthUseCase {

	return &AuthUseCase{
		userRepos:  user_repos,
		hashSalt:   hash_salt,
		signingKey: signing_key,
		expireTime: tokenTLSSeconds * 24 * time.Hour,
	}
}

func (a *AuthUseCase) HashPassword(p string) string {
	pwd := sha256.New()
	pwd.Write([]byte(p))
	pwd.Write([]byte(a.hashSalt))
	password := fmt.Sprintf("%x", pwd.Sum(nil))
	return password
}

func (a *AuthUseCase) SignUp(email, firstname, lastname, password string) error {

	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}

	user := &models.User{
		Id:        id.String(),
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  a.HashPassword(password),
	}

	return a.userRepos.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	User_id string `json:"user_id"`
}

func (a *AuthUseCase) GenerateAuthToken(user_id string) (string, error) {

	claims := tokenClaims{
		jwt.StandardClaims{
			Subject:   "authentification",
			ExpiresAt: int64(time.Now().Add(a.expireTime).Unix()),
			Issuer:    "auth-service",
		},
		user_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signed_str, err := token.SignedString(a.signingKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return signed_str, err
}

func (a *AuthUseCase) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(a.signingKey), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.User_id, nil
	} else {
		return "", errors.New("Error invalid access token")

	}
}

func (a *AuthUseCase) SignIn(email, password string) (string, string, error) {
	password = a.HashPassword(password)
	id, err := a.userRepos.GetUserID(email, password)
	if err != nil {
		log.Println(err)
		return "", "", errors.New("user not found")
	}

	accessToken, err := a.GenerateAuthToken(id)
	if err != nil {
		log.Println(err)
		return "", "", errors.New("Token generation error")
	}
	return accessToken, id, nil
}
