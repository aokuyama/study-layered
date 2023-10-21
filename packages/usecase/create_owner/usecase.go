package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type createOwner struct {
	factory    owner.OwnerFactory
	repository owner.OwnerRepository
}

type CreateOwnerInput struct {
}

type createOwnerOutput struct {
	Owner *owner.Owner
}

func New(f owner.OwnerFactory, r owner.OwnerRepository) *createOwner {
	u := createOwner{f, r}
	return &u
}

func (u *createOwner) Invoke(i *CreateOwnerInput) (*createOwnerOutput, error) {
	var err error
	o, err := u.factory.Create()
	if err != nil {
		return nil, err
	}
	err = u.repository.Save(o)
	if err != nil {
		return nil, err
	}
	out := createOwnerOutput{o}
	return &out, nil
}
