package user

type AuthToken = string

func NewAuthToken(s string) (*AuthToken, error) {
	t := AuthToken(s)
	return &t, nil
}
