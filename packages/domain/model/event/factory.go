//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package event

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type EventFactory interface {
	Create(circleID *circle.CircleID, name *string) (*EventEntity, error)
}

type EventFactoryImpl struct{}

func (f EventFactoryImpl) Create(circleID *circle.CircleID, name *string) (*EventEntity, error) {
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
	return &EventEntity{i, *circleID, *n, *p}, nil
}
