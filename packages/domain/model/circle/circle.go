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

type CircleInput struct {
	ID      string
	OwnerID string
	Name    string
	Path    string
}

func NewCircle(i *CircleInput) (*Circle, error) {
	ID, err := NewCircleID(i.ID)
	if err != nil {
		return nil, err
	}
	o, err := owner.NewOwnerID(i.OwnerID)
	if err != nil {
		return nil, err
	}
	n, err := NewName(i.Name)
	if err != nil {
		return nil, err
	}
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}

	c := Circle{*ID, *o, *n, *p}
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
