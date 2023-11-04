//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package user

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type UserFactory interface {
	Create() (*User, error)
}

type OwnerFactoryImpl struct{}

func (f OwnerFactoryImpl) Create() (*User, error) {
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	oi := UserID{i}
	return NewUser(&oi)
}
