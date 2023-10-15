//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

type CircleRepository interface {
	Create(*RegisterCircle) error
	Find(*CircleID) (*CircleEntity, error)
	FindByPath(*path.Path) (*CircleEntity, error)
}
