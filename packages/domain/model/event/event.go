package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type Event struct {
	id       EventID
	circleID circle.CircleID
	name     Name
	path     path.Path
	guest    GuestCollection
}

type EventInput struct {
	ID       string
	CircleID string
	Name     string
	Path     string
	Guest    []guest.GuestInput
}

func NewEvent(i *EventInput) (*Event, error) {
	ID, err := NewEventID(i.ID)
	if err != nil {
		return nil, err
	}
	c, err := circle.NewCircleID(i.CircleID)
	if err != nil {
		return nil, err
	}
	n, err := NewName(i.Name)
	if err != nil {
		return nil, err
	}
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}
	g, err := NewGuestCollection(i.Guest)
	if err != nil {
		return nil, err
	}

	e := Event{*ID, *c, *n, *p, *g}
	return &e, nil
}

func (e *Event) ID() *EventID {
	return &e.id
}
func (e *Event) CircleID() *circle.CircleID {
	return &e.circleID
}
func (e *Event) Name() *Name {
	return &e.name
}
func (e *Event) Path() *path.Path {
	return &e.path
}
func (e *Event) Guest() *GuestCollection {
	return &e.guest
}

func (en *Event) Identical(e *Event) bool {
	return en.ID().Equals(e.ID().UUID)
}

func (e *Event) JoinGuest(g *guest.Guest) *Event {
	newGuests := e.guest.AppendOrUpdate(g)
	if newGuests == nil {
		return nil
	}
	newEvent := Event{e.id, e.circleID, e.name, e.path, *newGuests}
	return &newEvent
}

func (e *Event) RemoveGuest(i *user.UserID) *Event {
	newGuests := e.guest.Remove(i)
	if newGuests == nil {
		return nil
	}
	newEvent := Event{e.id, e.circleID, e.name, e.path, *newGuests}
	return &newEvent
}
