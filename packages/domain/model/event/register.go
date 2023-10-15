package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type RegisterEvent struct {
	ID       *EventID
	CircleID *circle.CircleID
	Name     *Name
	Path     *path.Path
}

func newRegisterEvent(id *EventID, circleID *circle.CircleID, name *Name, path *path.Path) (*RegisterEvent, error) {
	e := RegisterEvent{id, circleID, name, path}
	return &e, nil
}

func GenerateRegisterEvent(circleID *circle.CircleID, name *string) (*RegisterEvent, error) {
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}
	u, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := EventID{u}
	p, err := path.GeneratePath()
	if err != nil {
		return nil, err
	}
	return newRegisterEvent(&i, circleID, n, p)
}
