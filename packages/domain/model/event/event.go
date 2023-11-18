package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

type Event struct {
	id       EventID
	circleID circle.CircleID
	name     Name
	path     path.Path
	guest    GuestCollection
}

func NewEvent(eventID *string, circleID *string, name *string, path *path.Path) (*Event, error) {
	i, err := NewEventID(*eventID)
	if err != nil {
		return nil, err
	}
	c, err := circle.NewCircleID(*circleID)
	if err != nil {
		return nil, err
	}
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}

	e := Event{*i, *c, *n, *path, *NewEmptyGuestCollection()}
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
