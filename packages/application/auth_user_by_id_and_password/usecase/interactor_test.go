package usecase_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/auth_user_by_id_and_password/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	mock_user "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user/.mock"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	id := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	pw := "passwordpassword"
	u := user.User{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_user.MockUserRepository)
		input           AuthUserByIDAndPasswordInput
		expect          *AuthUserByIDAndPasswordOutput
		expect_err      error
	}{
		{"success", func(r *mock_user.MockUserRepository) {
			r.EXPECT().FindWithPasswordAuth(gomock.Any(), gomock.Any()).Return(&u, nil)
		}, AuthUserByIDAndPasswordInput{id, pw}, &AuthUserByIDAndPasswordOutput{u}, nil},

		{"invalid id", func(r *mock_user.MockUserRepository) {
		}, AuthUserByIDAndPasswordInput{"", pw}, nil, errs.ErrBadParam},

		{"invalid password", func(r *mock_user.MockUserRepository) {
		}, AuthUserByIDAndPasswordInput{id, ""}, nil, errs.ErrBadParam},

		{"fail auth", func(r *mock_user.MockUserRepository) {
			r.EXPECT().FindWithPasswordAuth(gomock.Any(), gomock.Any()).Return(nil, errs.NewUnauthorized("test"))
		}, AuthUserByIDAndPasswordInput{id, pw}, nil, errs.ErrUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mock_user.NewMockUserRepository(ctrl)
			tt.mock_repository(r)

			u := New(r)
			o, err := u.Invoke(&tt.input)
			assert.Equal(t, tt.expect, o)
			assert.True(t, errors.Is(err, tt.expect_err))
		})
	}
}
