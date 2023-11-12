package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type EventEntity struct {
	id       EventID
	circleID circle.CircleID
	name     Name
	path     path.Path
	guest    GuestCollection
}

func NewEventEntity(id *string, circleID *string, name *string, path *path.Path) (*EventEntity, error) {
	i, err := NewEventID(*id)
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

	e := EventEntity{*i, *c, *n, *path, *NewEmptyGuestCollection()}
	return &e, nil
}

func (e *EventEntity) ID() *EventID {
	return &e.id
}
func (e *EventEntity) CircleID() *circle.CircleID {
	return &e.circleID
}
func (e *EventEntity) Name() *Name {
	return &e.name
}
func (e *EventEntity) Path() *path.Path {
	return &e.path
}
func (e *EventEntity) Guest() *GuestCollection {
	return &e.guest
}

func (en *EventEntity) Identical(e *EventEntity) bool {
	return en.ID().Equals(e.ID().UUID)
}
