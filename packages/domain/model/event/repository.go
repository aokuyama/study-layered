//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package event

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

type EventRepository interface {
	Create(*RegisterEvent) error
	FindByPath(*path.Path) (*EventEntity, error)
}
