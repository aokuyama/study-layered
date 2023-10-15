package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type createCircle struct {
	ownerRepository  owner.OwnerRepository
	circleRepository circle.CircleRepository
}

type CreateCircleInput struct {
	OwnerID    string
	CreateName string
}

type createCircleOutput struct {
	Circle *circle.Circle
}

func New(o owner.OwnerRepository, c circle.CircleRepository) *createCircle {
	u := createCircle{o, c}
	return &u
}

func (u *createCircle) Invoke(i *CreateCircleInput) (*createCircleOutput, error) {
	var c *circle.Circle
	var err error
	id, err := common.NewUUID(i.OwnerID)
	if err != nil {
		return nil, err
	}

	_, err = u.ownerRepository.Find(id)
	if err != nil {
		return nil, err
	}

	c, err = circle.GenerateCircle(&i.CreateName)
	if err != nil {
		return nil, err
	}
	c, err = u.circleRepository.Save(c)
	if err != nil {
		return nil, err
	}
	o := createCircleOutput{c}
	return &o, nil
}
