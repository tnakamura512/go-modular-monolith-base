package umodel

type UserRepositroyInterface interface {
	SaveUser(token, name string) error
	GetIdByToken(token string) (int, error)
}
