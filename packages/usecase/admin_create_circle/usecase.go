package admin_create_circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
)

type Usecase struct {
	ownerRepository  owner.OwnerRepository
	circleRepository circle.CircleRepository
}

type Input struct {
	OwnerID string
	Name    string
}

type Output struct {
	Circle *circle.Circle
}

func New(o owner.OwnerRepository, c circle.CircleRepository) *Usecase {
	u := Usecase{o, c}
	return &u
}

func (u *Usecase) Invoke(i *Input) (*Output, error) {
	var c *circle.Circle
	var err error
	id, err := common.NewUUID(i.OwnerID)
	if err != nil {
		return nil, err
	}

	_, err = u.ownerRepository.LoadByID(id)
	if err != nil {
		return nil, err
	}

	c, err = circle.GenerateCircle(&i.Name)
	if err != nil {
		return nil, err
	}
	c, err = u.circleRepository.Save(c)
	if err != nil {
		return nil, err
	}
	o := Output{c}
	return &o, nil
}
