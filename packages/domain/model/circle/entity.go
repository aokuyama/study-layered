package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type CircleEntity struct {
	ID      *CircleID
	OwnerID *owner.OwnerID
	Name    *Name
}

func NewCircleEntity(id *CircleID, ownerID *owner.OwnerID, name *Name) (*CircleEntity, error) {
	c := CircleEntity{id, ownerID, name}
	return &c, nil
}

func (e *CircleEntity) Identical(c *CircleEntity) bool {
	return e.ID.Equals(c.ID)
}
