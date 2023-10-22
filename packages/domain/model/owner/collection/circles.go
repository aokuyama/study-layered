package collection

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

type OwnerCircles struct {
	circles []circle.CircleID
}

func NewOwnerCircles(circles []circle.CircleID) *OwnerCircles {
	c := OwnerCircles{circles}
	return &c
}

func (c *OwnerCircles) IsAppendable() bool {
	return len(c.circles) < 1
}
