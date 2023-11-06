package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

type AuthUserByIDAndPasswordUsecase interface {
	Invoke(i *AuthUserByIDAndPasswordInput) (*AuthUserByIDAndPasswordOutput, error)
}

type AuthUserByIDAndPasswordInput struct {
	ID       string
	Password string
}

type AuthUserByIDAndPasswordOutput struct {
	User user.User
}
