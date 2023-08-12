package circle

type Circle struct {
	ID CircleID
}

func NewCircle(id *CircleID) (*Circle, error) {
	c := Circle{*id}
	return &c, nil
}

func (e *Circle) Identical(c *Circle) bool {
	return e.ID.Equals(&c.ID)
}
