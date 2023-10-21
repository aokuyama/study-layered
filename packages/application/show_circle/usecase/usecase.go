package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type showCircle struct {
	circleRepository circle.CircleRepository
}

type ShowCircleInput struct {
	Path string
}

type showCircleOutput struct {
	Circle *circle.CircleEntity
}

func New(c circle.CircleRepository) *showCircle {
	u := showCircle{c}
	return &u
}

func (u *showCircle) Invoke(i *ShowCircleInput) (*showCircleOutput, error) {
	var err error
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}

	c, err := u.circleRepository.FindByPath(p)
	if err != nil {
		return nil, err
	}

	o := showCircleOutput{c}
	return &o, nil
}
