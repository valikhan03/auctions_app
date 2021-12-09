package auth

type UseCase interface {
	SignUp(email, firstname, lastname, password string) error
	SignIn(email, password string) (string, string, error)
	ParseToken(accessToken string) (string, error)
	GenerateAuthToken(user_id string) (string, error)
	HashPassword(p string) string
}
