package owner

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type OwnerID struct {
	*common.UUID
}

func NewOwnerID(v string) (*OwnerID, error) {
	u, err := common.NewUUID(v)
	if err != nil {
		return nil, err
	}
	i := OwnerID{u}
	return &i, err
}
