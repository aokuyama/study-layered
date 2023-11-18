package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

type UserJoinToEventUsecase interface {
	Invoke(i *UserJoinToEventInput) (*UserJoinToEventOutput, error)
}

type UserJoinToEventInput struct {
	EventID string
	UserID  string
	Name    string
	Number  uint8
}

type UserJoinToEventOutput struct {
	Event event.EventEntity
	Guest guest.Guest
}
