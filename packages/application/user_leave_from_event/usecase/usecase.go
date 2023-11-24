package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type UserLeaveFromEventUsecase interface {
	Invoke(i *UserLeaveFromEventInput) (*UserLeaveFromEventOutput, error)
}

type UserLeaveFromEventInput struct {
	EventID string
	UserID  string
}

type UserLeaveFromEventOutput struct {
	Event event.Event
}
