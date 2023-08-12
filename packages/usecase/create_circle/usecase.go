package create_circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

type Usecase struct {
	repository circle.CircleRepository
}

type Input struct {
	name string
}

type Output struct {
	circle *circle.Circle
}

func (u *Usecase) Invoke(i *Input) (*Output, error) {
	var c *circle.Circle
	var err error
	c, err = circle.GenerateCircle(&i.name)
	if err != nil {
		return nil, err
	}
	c, err = u.repository.Save(c)
	if err != nil {
		return nil, err
	}
	o := Output{c}
	return &o, nil
}
