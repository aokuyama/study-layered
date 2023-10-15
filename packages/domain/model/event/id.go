package event

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type EventID = common.UUID

func NewEventID(v string) (*EventID, error) {
	return common.NewUUID(v)
}
