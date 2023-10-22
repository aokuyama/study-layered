package usecase

import (
	"errors"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/collection"
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

	cl, err := u.circleRepository.SearchByOwner(ownerID)
	if err != nil {
		panic(err)
	}
	cc := collection.NewOwnerCircles(*cl)
	if !cc.IsAppendable() {
		return nil, errors.New("unable to append circle")
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
