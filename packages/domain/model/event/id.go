package event

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type EventID struct {
	*common.UUID
}

func NewEventID(v string) (*EventID, error) {
	u, err := common.NewUUID(v)
	if err != nil {
		return nil, err
	}
	i := EventID{u}
	return &i, err
}
