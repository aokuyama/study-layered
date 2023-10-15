//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

type CircleRepository interface {
	Save(*Circle) error
	Find(*CircleID) (*Circle, error)
	FindByPath(*path.Path) (*Circle, error)
}
