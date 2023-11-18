//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type CircleRepository interface {
	Create(*Circle) error
	Find(*CircleID) (*Circle, error)
	FindByPath(*path.Path) (*Circle, error)
	SearchByOwner(*owner.OwnerID) (*[]CircleID, error)
}
