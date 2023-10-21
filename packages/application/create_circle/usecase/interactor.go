package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type createCircleInteractor struct {
	circleFactory    circle.CircleFactory
	ownerRepository  owner.OwnerRepository
	circleRepository circle.CircleRepository
}

func New(f circle.CircleFactory, or owner.OwnerRepository, cr circle.CircleRepository) CreateCircleUsecase {
	u := createCircleInteractor{f, or, cr}
	return &u
}

func (u *createCircleInteractor) Invoke(i *CreateCircleInput) (*CreateCircleOutput, error) {
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
	o := CreateCircleOutput{c}
	return &o, nil
}
