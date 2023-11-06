package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type authUserByIDAndPasswordInteractor struct {
	repository user.UserRepository
}

func New(r user.UserRepository) AuthUserByIDAndPasswordUsecase {
	u := authUserByIDAndPasswordInteractor{r}
	return &u
}

func (u *authUserByIDAndPasswordInteractor) Invoke(i *AuthUserByIDAndPasswordInput) (*AuthUserByIDAndPasswordOutput, error) {
	var err error
	id, err := user.NewUserID(i.ID)
	if err != nil {
		return nil, err
	}
	pw, err := user.NewPassword(i.Password)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.FindWithPasswordAuth(id, pw)
	if err != nil {
		return nil, err
	}
	out := AuthUserByIDAndPasswordOutput{*user}
	return &out, nil
}
