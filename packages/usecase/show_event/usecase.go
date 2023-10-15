package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type showEvent struct {
	eventRepository event.EventRepository
}

type ShowEventInput struct {
	Path string
}

type showEventOutput struct {
	Event *event.EventEntity
}

func New(c event.EventRepository) *showEvent {
	u := showEvent{c}
	return &u
}

func (u *showEvent) Invoke(i *ShowEventInput) (*showEventOutput, error) {
	var err error
	p, err := path.NewPath(i.Path)
	if err != nil {
		return nil, err
	}

	e, err := u.eventRepository.FindByPath(p)
	if err != nil {
		return nil, err
	}

	o := showEventOutput{e}
	return &o, nil
}
