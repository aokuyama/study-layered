package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

type GuestCollection struct {
	guest []guest.Guest
}

func NewEmptyGuestCollection() *GuestCollection {
	g := []guest.Guest{}
	c := GuestCollection{g}
	return &c
}

func (c *GuestCollection) IsEmpty() bool {
	return len(c.guest) == 0
}
