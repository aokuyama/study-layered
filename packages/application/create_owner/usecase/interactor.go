package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type createOwnerInteractor struct {
	factory    owner.OwnerFactory
	repository owner.OwnerRepository
}

func New(f owner.OwnerFactory, r owner.OwnerRepository) CreateOwnerUsecase {
	u := createOwnerInteractor{f, r}
	return &u
}

func (u *createOwnerInteractor) Invoke(i *CreateOwnerInput) (*CreateOwnerOutput, error) {
	var err error
	o, err := u.factory.Create()
	if err != nil {
		return nil, err
	}
	err = u.repository.Save(o)
	if err != nil {
		return nil, err
	}
	out := CreateOwnerOutput{o}
	return &out, nil
}
