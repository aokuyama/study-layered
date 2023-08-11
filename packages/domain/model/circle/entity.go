package circle

import (
	"aokuyama/circle_scheduler-api/packages/domain/model/common"
)

type Circle struct {
	ID common.UUID
}

func NewCircle(id common.UUID) (*Circle, error) {
	c := Circle{id}
	return &c, nil
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID.Equals(&c.ID)
}
