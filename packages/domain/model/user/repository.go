//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package user

type UserRepository interface {
	Create(*UserWithPassword) error
}

type UserAuthRepository interface {
	CreateToken(*UserID) (*AuthToken, error)
	AuthByToken(*UserID, *AuthToken) error
}
