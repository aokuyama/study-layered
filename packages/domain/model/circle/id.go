package circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type CircleID struct {
	*common.UUID
}

func NewCircleID(v string) (*CircleID, error) {
	u, err := common.NewUUID(v)
	if err != nil {
		return nil, err
	}
	i := CircleID{u}
	return &i, err
}
