package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
)

type EventEntity struct {
	ID       *EventID
	CircleID *circle.CircleID
	Name     *Name
}

func NewEventEntity(id *EventID, circleID *circle.CircleID, name *Name) (*EventEntity, error) {
	e := EventEntity{id, circleID, name}
	return &e, nil
}

func (e *EventEntity) Identical(c *EventEntity) bool {
	return e.ID.Equals(c.ID)
}
