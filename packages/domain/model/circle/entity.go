package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type Circle struct {
	ID      *CircleID
	OwnerID *owner.OwnerID
	Name    *Name
	Path    *path.Path
}

func NewCircle(id *CircleID, ownerID *owner.OwnerID, name *Name, path *path.Path) (*Circle, error) {
	c := Circle{id, ownerID, name, path}
	return &c, nil
}

func GenerateCircle(ownerID *owner.OwnerID, name *string) (*Circle, error) {
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	p, err := path.GeneratePath()
	if err != nil {
		return nil, err
	}
	return NewCircle(i, ownerID, n, p)
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID.Equals(c.ID)
}
