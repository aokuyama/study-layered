package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type createEvent struct {
	circleRepository circle.CircleRepository
	eventRepository  event.EventRepository
}

type CreateEventInput struct {
	CircleID  string
	EventName string
}

type createEventOutput struct {
	Event *event.RegisterEvent
}

func New(o circle.CircleRepository, c event.EventRepository) *createEvent {
	u := createEvent{o, c}
	return &u
}

func (u *createEvent) Invoke(i *CreateEventInput) (*createEventOutput, error) {
	var err error
	circleID, err := circle.NewCircleID(i.CircleID)
	if err != nil {
		return nil, err
	}

	_, err = u.circleRepository.Find(circleID)
	if err != nil {
		return nil, err
	}

	e, err := event.GenerateRegisterEvent(circleID, &i.EventName)
	if err != nil {
		return nil, err
	}
	err = u.eventRepository.Create(e)
	if err != nil {
		return nil, err
	}
	o := createEventOutput{e}
	return &o, nil
}
