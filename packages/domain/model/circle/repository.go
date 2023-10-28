//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type CircleRepository interface {
	Create(*CircleEntity) error
	Find(*CircleID) (*CircleEntity, error)
	FindByPath(*path.Path) (*CircleEntity, error)
	SearchByOwner(*owner.OwnerID) (*[]CircleID, error)
}
