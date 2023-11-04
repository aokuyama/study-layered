package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type createUserByPasswordInteractor struct {
	factory    user.UserFactory
	repository user.UserRepository
}

func New(f user.UserFactory, r user.UserRepository) CreateUserByPasswordUsecase {
	u := createUserByPasswordInteractor{f, r}
	return &u
}

func (u *createUserByPasswordInteractor) Invoke(i *CreateUserByPasswordInput) (*CreateUserByPasswordOutput, error) {
	var err error
	user, err := u.factory.Create(i.Password)
	if err != nil {
		return nil, err
	}
	err = u.repository.Create(user)
	if err != nil {
		return nil, err
	}
	out := CreateUserByPasswordOutput{user.User}
	return &out, nil
}
