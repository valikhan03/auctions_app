package auth

type UseCase interface {
	SignUp(email, password, firstname, lastname string) error
	SignIn(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	GenerateAuthToken(user_id string) (string, error)
}
