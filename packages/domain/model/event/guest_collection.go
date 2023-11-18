package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

type GuestCollection struct {
	list []guest.Guest
}

func NewGuestCollection() *GuestCollection {
	g := []guest.Guest{}
	c := GuestCollection{g}
	return &c
}

func NewEmptyGuestCollection() *GuestCollection {
	g := []guest.Guest{}
	c := GuestCollection{g}
	return &c
}

func (c *GuestCollection) Empty() bool {
	return len(c.list) == 0
}

func (c *GuestCollection) Len() int {
	return len(c.list)
}

func (c *GuestCollection) Nth(n int) *guest.Guest {
	return &c.list[n]
}

func (c *GuestCollection) Append(g *guest.Guest) *GuestCollection {
	l := []guest.Guest{}
	for _, item := range c.list {
		if item.Identical(g) {
			return nil
		}
		l = append(l, item)
	}
	l = append(l, *g)
	new := GuestCollection{l}
	return &new
}

func (c *GuestCollection) Update(g *guest.Guest) *GuestCollection {
	l := []guest.Guest{}
	found := false
	for _, item := range c.list {
		if item.Identical(g) {
			if item.Equals(g) {
				return nil
			}
			found = true
			l = append(l, *g)
		} else {
			l = append(l, item)
		}
	}
	if !found {
		return nil
	}
	new := GuestCollection{l}
	return &new
}

func (c *GuestCollection) AppendOrUpdate(g *guest.Guest) *GuestCollection {
	l := []guest.Guest{}
	found := false
	for _, item := range c.list {
		if item.Identical(g) {
			if item.Equals(g) {
				return nil
			}
			found = true
			l = append(l, *g)
		} else {
			l = append(l, item)
		}
	}
	if !found {
		l = append(l, *g)
	}
	new := GuestCollection{l}
	return &new
}
