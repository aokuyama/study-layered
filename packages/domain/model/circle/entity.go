package circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type Circle struct {
	ID   *CircleID
	Name *Name
}

func NewCircle(id *CircleID, name *Name) (*Circle, error) {
	c := Circle{id, name}
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
	c := Circle{i, n}
	return &c, nil
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID.Equals(c.ID)
}
