package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type createCircle struct {
	circleFactory    circle.CircleFactory
	ownerRepository  owner.OwnerRepository
	circleRepository circle.CircleRepository
}

type CreateCircleInput struct {
	OwnerID    string
	CircleName string
}

type createCircleOutput struct {
	Circle *circle.RegisterCircle
}

func New(f circle.CircleFactory, or owner.OwnerRepository, cr circle.CircleRepository) *createCircle {
	u := createCircle{f, or, cr}
	return &u
}

func (u *createCircle) Invoke(i *CreateCircleInput) (*createCircleOutput, error) {
	var err error
	ownerID, err := owner.NewOwnerID(i.OwnerID)
	if err != nil {
		return nil, err
	}

	_, err = u.ownerRepository.Find(ownerID)
	if err != nil {
		return nil, err
	}

	c, err := u.circleFactory.Create(ownerID, &i.CircleName)
	if err != nil {
		return nil, err
	}
	err = u.circleRepository.Create(c)
	if err != nil {
		return nil, err
	}
	o := createCircleOutput{c}
	return &o, nil
}
