//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package owner

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type OwnerFactory interface {
	Create() (*Owner, error)
}

type OwnerFactoryImpl struct{}

func (f OwnerFactoryImpl) Create() (*Owner, error) {
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	oi := OwnerID{i}
	return NewOwner(&oi)
}
