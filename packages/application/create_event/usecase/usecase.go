package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type CreateEventUsecase interface {
	Invoke(i *CreateEventInput) (*CreateEventOutput, error)
}

type CreateEventInput struct {
	CircleID  string
	EventName string
}

type CreateEventOutput struct {
	Event *event.RegisterEvent
}
