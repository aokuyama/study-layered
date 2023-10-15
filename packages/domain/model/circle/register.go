package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
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

func GenerateRegisterCircle(ownerID *owner.OwnerID, name *string) (*RegisterCircle, error) {
	n, err := NewName(*name)
	if err != nil {
		return nil, err
	}
	u, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := CircleID{u}
	p, err := path.GeneratePath()
	if err != nil {
		return nil, err
	}
	return newRegisterCircle(&i, ownerID, n, p)
}
