package usecase

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

type UserCreateAuthTokenUsecase interface {
	Invoke(i *UserCreateAuthTokenInput) (*UserCreateAuthTokenOutput, error)
}

type UserCreateAuthTokenInput struct {
	UserId string
}

type UserCreateAuthTokenOutput struct {
	Token user.AuthToken
}
