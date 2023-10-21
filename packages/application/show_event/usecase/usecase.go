package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

type ShowEventUsecase interface {
	Invoke(i *ShowEventInput) (*ShowEventOutput, error)
}

type ShowEventInput struct {
	Path string
}

type ShowEventOutput struct {
	Event *event.EventEntity
}
