//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package user

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type UserFactory interface {
	Create(password string) (*UserWithPassword, error)
}

type UserFactoryImpl struct{}

func (f UserFactoryImpl) Create(password string) (*UserWithPassword, error) {
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	ui := UserID{i}
	user, err := NewUser(&ui)
	if err != nil {
		return nil, err
	}
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	u := UserWithPassword{User: *user, Password: *p}
	return &u, nil
}
