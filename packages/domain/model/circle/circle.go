package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type Circle struct {
	id      CircleID
	ownerID owner.OwnerID
	name    Name
	path    path.Path
}

func NewCircle(id *string, ownerID *string, name *string, path *path.Path) (*Circle, error) {
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

	c := Circle{*i, *o, *n, *path}
	return &c, nil
}

func (e *Circle) ID() *CircleID {
	return &e.id
}
func (e *Circle) OwnerID() *owner.OwnerID {
	return &e.ownerID
}
func (e *Circle) Name() *Name {
	return &e.name
}
func (e *Circle) Path() *path.Path {
	return &e.path
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID().Equals(c.ID().UUID)
}
