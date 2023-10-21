package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type CreateOwnerUsecase interface {
	Invoke(i *CreateOwnerInput) (*CreateOwnerOutput, error)
}

type CreateOwnerInput struct {
}

type CreateOwnerOutput struct {
	Owner *owner.Owner
}
