package registry

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

type Dummy struct {
}

func (r *Dummy) Save(c *circle.Circle) (*circle.Circle, error) {
	return c, nil
}

func (r *Dummy) LoadByID(*circle.CircleID) (*circle.Circle, error) {
	n := "dummy"
	c, err := circle.GenerateCircle(&n)
	return c, err
}
