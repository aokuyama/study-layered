package get_circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
)

type Usecase struct {
	repository circle.CircleRepository
}

type Input struct {
	id string
}

type Output struct {
	circle *circle.Circle
}

func (u *Usecase) Invoke(i *Input) (*Output, error) {
	var err error
	id, err := common.NewUUID(i.id)
	if err != nil {
		return nil, err
	}
	c, err := u.repository.LoadByID(id)
	if err != nil {
		return nil, err
	}
	o := Output{c}
	return &o, nil
}
