package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type userCreateAuthTokenInteractor struct {
	repository user.UserAuthRepository
}

func New(r user.UserAuthRepository) UserCreateAuthTokenUsecase {
	u := userCreateAuthTokenInteractor{r}
	return &u
}

func (u *userCreateAuthTokenInteractor) Invoke(i *UserCreateAuthTokenInput) (*UserCreateAuthTokenOutput, error) {
	var err error
	id, err := user.NewUserID(i.UserId)
	if err != nil {
		return nil, err
	}
	t, err := u.repository.CreateToken(id)
	if err != nil {
		return nil, err
	}
	out := UserCreateAuthTokenOutput{*t}
	return &out, nil
}
