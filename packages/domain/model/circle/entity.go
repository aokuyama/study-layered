package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type CircleEntity struct {
	id      CircleID
	ownerID owner.OwnerID
	name    Name
}

func NewCircleEntity(id *string, ownerID *string, name *string) (*CircleEntity, error) {
	i, err := NewCircleID(*id)
	if err != nil {
		return nil, err
	}
	o, err := owner.NewOwnerID(*ownerID)
	if err != nil {
		return nil, err
	}
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}

	c := CircleEntity{*i, *o, *n}
	return &c, nil
}

func (e *CircleEntity) ID() *CircleID {
	return &e.id
}
func (e *CircleEntity) OwnerID() *owner.OwnerID {
	return &e.ownerID
}
func (e *CircleEntity) Name() *Name {
	return &e.name
}

func (e *CircleEntity) Identical(c *CircleEntity) bool {
	return e.ID().Equals(c.ID().UUID)
}
