//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package owner

type OwnerRepository interface {
	Save(*Owner) (*Owner, error)
	Find(*OwnerID) (*Owner, error)
}
