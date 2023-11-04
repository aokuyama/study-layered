//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package user

type UserRepository interface {
	Save(*User) error
	Find(*UserID) (*User, error)
}
