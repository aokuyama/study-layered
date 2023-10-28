package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
)

type CreateCircleUsecase interface {
	Invoke(i *CreateCircleInput) (*CreateCircleOutput, error)
}

type CreateCircleInput struct {
	OwnerID    string
	CircleName string
}

type CreateCircleOutput struct {
	Circle *circle.CircleEntity
}
