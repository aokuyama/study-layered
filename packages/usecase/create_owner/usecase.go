package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type createOwner struct {
	repository owner.OwnerRepository
}

type CreateOwnerInput struct {
}

type createOwnerOutput struct {
	Owner *owner.Owner
}

func New(r owner.OwnerRepository) *createOwner {
	u := createOwner{r}
	return &u
}

func (u *createOwner) Invoke(i *CreateOwnerInput) (*createOwnerOutput, error) {
	var err error
	o, err := owner.GenerateOwner()
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
