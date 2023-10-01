package owner

type OwnerRepository interface {
	Save(*Owner) (*Owner, error)
	LoadByID(*OwnerID) (*Owner, error)
}
