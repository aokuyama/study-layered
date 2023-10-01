package create_circle

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

type Usecase struct {
	repository owner.OwnerRepository
}

type Input struct {
}

type Output struct {
	Owner *owner.Owner
}

func New(r owner.OwnerRepository) *Usecase {
	u := Usecase{r}
	return &u
}

func (u *Usecase) Invoke(i *Input) (*Output, error) {
	var o *owner.Owner
	var err error
	o, err = owner.GenerateOwner()
	if err != nil {
		return nil, err
	}
	o, err = u.repository.Save(o)
	if err != nil {
		return nil, err
	}
	out := Output{o}
	return &out, nil
}
