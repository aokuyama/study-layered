package owner

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type OwnerID = common.UUID

func NewOwnerID(v string) (*OwnerID, error) {
	return common.NewUUID(v)
}
