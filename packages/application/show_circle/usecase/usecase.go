package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
)

type ShowCircleUsecase interface {
	Invoke(i *ShowCircleInput) (*ShowCircleOutput, error)
}

type ShowCircleInput struct {
	Path string
}

type ShowCircleOutput struct {
	Circle *circle.Circle
}
