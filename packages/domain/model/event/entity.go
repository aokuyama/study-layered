package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
)

type EventEntity struct {
	ID       *EventID
	CircleID *circle.CircleID
	Name     *Name
}

func NewEventEntity(id *string, circleID *string, name *string) (*EventEntity, error) {
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

	e := EventEntity{i, c, n}
	return &e, nil
}

func (en *EventEntity) Identical(e *EventEntity) bool {
	return en.ID.Equals(e.ID.UUID)
}
