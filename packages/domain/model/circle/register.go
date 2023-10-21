package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type RegisterCircle struct {
	ID      *CircleID
	OwnerID *owner.OwnerID
	Name    *Name
	Path    *path.Path
}

func newRegisterCircle(id *CircleID, ownerID *owner.OwnerID, name *Name, path *path.Path) (*RegisterCircle, error) {
	c := RegisterCircle{id, ownerID, name, path}
	return &c, nil
}
