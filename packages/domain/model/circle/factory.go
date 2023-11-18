//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type CircleFactory interface {
	Create(ownerID *owner.OwnerID, name *string) (*Circle, error)
}

type CircleFactoryImpl struct{}

func (f CircleFactoryImpl) Create(ownerID *owner.OwnerID, name *string) (*Circle, error) {
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
	return &Circle{i, *ownerID, *n, *p}, nil
}
