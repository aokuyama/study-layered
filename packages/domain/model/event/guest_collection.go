package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type GuestCollection struct {
	list []guest.Guest
}

func NewGuestCollection(i []guest.GuestInput) (*GuestCollection, error) {
	c := NewEmptyGuestCollection()
	for _, v := range i {
		g, err := guest.NewGuest(&v)
		if err != nil {
			return nil, err
		}
		c = c.Append(g)
		if c == nil {
			return nil, errs.NewFatal("duplication entity")
		}
	}
	return c, nil
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

func (c *GuestCollection) Items() []guest.Guest {
	return c.list
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

func (c *GuestCollection) IdenticalItem(g *guest.Guest) *guest.Guest {
	for _, item := range c.list {
		if item.Identical(g) {
			return &item
		}
	}
	return nil
}

func (c *GuestCollection) ExistsIdentical(g *guest.Guest) bool {
	return c.IdenticalItem(g) != nil
}

func (c *GuestCollection) Remove(id *user.UserID) *GuestCollection {
	l := []guest.Guest{}
	found := false
	for _, item := range c.list {
		if item.UserID().Equals(id.UUID) {
			found = true
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
