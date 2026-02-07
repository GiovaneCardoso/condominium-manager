package authrepository

type AuthRepository interface {
	Authenticate(email, password string) (string, error)
}
