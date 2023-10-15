package circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type CircleID = common.UUID

func NewCircleID(v string) (*CircleID, error) {
	return common.NewUUID(v)
}
