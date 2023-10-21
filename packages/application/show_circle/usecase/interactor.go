package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
)

type showCircleInteractor struct {
	circleRepository circle.CircleRepository
}

func New(c circle.CircleRepository) ShowCircleUsecase {
	u := showCircleInteractor{c}
	return &u
}

func (u *showCircleInteractor) Invoke(i *ShowCircleInput) (*ShowCircleOutput, error) {
	var err error
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}

	c, err := u.circleRepository.FindByPath(p)
	if err != nil {
		return nil, err
	}

	o := ShowCircleOutput{c}
	return &o, nil
}
