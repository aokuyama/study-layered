package get_circle

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
)

type Usecase struct {
	repository circle.CircleRepository
}

func New(r circle.CircleRepository) *Usecase {
	u := Usecase{r}
	return &u
}

type Input struct {
	ID string
}

type Output struct {
	Circle *circle.Circle
}

func (u *Usecase) Invoke(i *Input) (*Output, error) {
	var err error
	id, err := common.NewUUID(i.ID)
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
