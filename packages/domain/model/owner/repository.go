package owner

type OwnerRepository interface {
	Save(*Owner) (*Owner, error)
	LoadByID(*Owner) (*Owner, error)
}
