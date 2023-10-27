package owner

type Owner struct {
	id OwnerID
}

func NewOwner(id *OwnerID) (*Owner, error) {
	c := Owner{*id}
	return &c, nil
}

func (e *Owner) ID() *OwnerID {
	return &e.id
}

func (e *Owner) Identical(c *Owner) bool {
	return e.ID().Equals(c.ID().UUID)
}
