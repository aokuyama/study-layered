package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type createEventInteractor struct {
	eventFactory     event.EventFactory
	circleRepository circle.CircleRepository
	eventRepository  event.EventRepository
}

func New(f event.EventFactory, or circle.CircleRepository, cr event.EventRepository) CreateEventUsecase {
	u := createEventInteractor{f, or, cr}
	return &u
}

func (u *createEventInteractor) Invoke(i *CreateEventInput) (*CreateEventOutput, error) {
	var err error
	circleID, err := circle.NewCircleID(i.CircleID)
	if err != nil {
		return nil, err
	}

	_, err = u.circleRepository.Find(circleID)
	if err != nil {
		return nil, err
	}

	e, err := u.eventFactory.Create(circleID, &i.EventName)
	if err != nil {
		return nil, err
	}
	err = u.eventRepository.Create(e)
	if err != nil {
		return nil, err
	}
	o := CreateEventOutput{e}
	return &o, nil
}
