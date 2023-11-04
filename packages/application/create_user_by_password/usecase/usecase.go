package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

type CreateUserByPasswordUsecase interface {
	Invoke(i *CreateUserByPasswordInput) (*CreateUserByPasswordOutput, error)
}

type CreateUserByPasswordInput struct {
	Password string
}

type CreateUserByPasswordOutput struct {
	User user.User
}
