package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type Event struct {
	ID       *EventID
	CircleID *circle.CircleID
	Name     *Name
	Path     *path.Path
}

func NewEvent(id *EventID, circleID *circle.CircleID, name *Name, path *path.Path) (*Event, error) {
	c := Event{id, circleID, name, path}
	return &c, nil
}

func GenerateEvent(circleID *circle.CircleID, name *string) (*Event, error) {
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	p, err := path.GeneratePath()
	if err != nil {
		return nil, err
	}
	return NewEvent(i, circleID, n, p)
}

func (e *Event) Identical(c *Event) bool {
	return e.ID.Equals(c.ID)
}
