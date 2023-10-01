package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type Circle struct {
	ID   *CircleID
	Name *Name
	Path *path.Path
}

func NewCircle(id *CircleID, name *Name, path *path.Path) (*Circle, error) {
	c := Circle{id, name, path}
	return &c, nil
}

func GenerateCircle(name *string) (*Circle, error) {
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
	return NewCircle(i, n, p)
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID.Equals(c.ID)
}
