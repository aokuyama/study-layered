package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type showEventInteractor struct {
	eventRepository event.EventRepository
}

func New(c event.EventRepository) ShowEventUsecase {
	u := showEventInteractor{c}
	return &u
}

func (u *showEventInteractor) Invoke(i *ShowEventInput) (*ShowEventOutput, error) {
	var err error
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}

	e, err := u.eventRepository.FindByPath(p)
	if err != nil {
		return nil, err
	}

	o := ShowEventOutput{e}
	return &o, nil
}
