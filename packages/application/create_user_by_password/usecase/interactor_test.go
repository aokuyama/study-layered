package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/create_user_by_password/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	mock_user "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user/.mock"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	u := user.User{}
	p, _ := user.NewPassword("passwordpassword")
	up := user.UserWithPassword{User: u, Password: *p}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_factory    func(r *mock_user.MockUserFactory)
		mock_repository func(r *mock_user.MockUserRepository)
		input           CreateUserByPasswordInput
		expect          *CreateUserByPasswordOutput
		expect_err      error
	}{
		{"success", func(r *mock_user.MockUserFactory) {
			r.EXPECT().Create("i1").Return(&up, nil)
		}, func(r *mock_user.MockUserRepository) {
			r.EXPECT().Create(&up).Return(nil)
		}, CreateUserByPasswordInput{"i1"}, &CreateUserByPasswordOutput{u}, nil},

		{"factory fail create", func(r *mock_user.MockUserFactory) {
			r.EXPECT().Create("i2").Return(nil, fmt.Errorf("test: %w", errs.ErrBadParam))
		}, func(r *mock_user.MockUserRepository) {
		}, CreateUserByPasswordInput{"i2"}, nil, errs.ErrBadParam},

		{"repository fail create", func(r *mock_user.MockUserFactory) {
			r.EXPECT().Create("i3").Return(&up, nil)
		}, func(r *mock_user.MockUserRepository) {
			r.EXPECT().Create(&up).Return(fmt.Errorf("test: %w", errs.ErrFatal))
		}, CreateUserByPasswordInput{"i3"}, nil, errs.ErrFatal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := mock_user.NewMockUserFactory(ctrl)
			tt.mock_factory(f)
			r := mock_user.NewMockUserRepository(ctrl)
			tt.mock_repository(r)

			u := New(f, r)
			o, err := u.Invoke(&tt.input)
			assert.Equal(t, tt.expect, o)
			assert.True(t, errors.Is(err, tt.expect_err))
		})
	}
}
