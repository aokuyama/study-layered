package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/user_create_auth_token/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	mock_user "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user/.mock"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	to, _ := user.NewAuthToken("a")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_user.MockUserAuthRepository)
		input           UserCreateAuthTokenInput
		expect          *UserCreateAuthTokenOutput
		expect_err      error
	}{
		{"success", func(r *mock_user.MockUserAuthRepository) {
			r.EXPECT().CreateToken(gomock.Any()).Return(to, nil)
		}, UserCreateAuthTokenInput{"26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"}, &UserCreateAuthTokenOutput{*to}, nil},

		{"invalid uuid", func(r *mock_user.MockUserAuthRepository) {
		}, UserCreateAuthTokenInput{"fail"}, nil, errs.ErrBadParam},

		{"fail create token", func(r *mock_user.MockUserAuthRepository) {
			r.EXPECT().CreateToken(gomock.Any()).Return(to, fmt.Errorf("test: %w", errs.ErrFatal))
		}, UserCreateAuthTokenInput{"26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"}, nil, errs.ErrFatal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mock_user.NewMockUserAuthRepository(ctrl)
			tt.mock_repository(r)

			u := New(r)
			o, err := u.Invoke(&tt.input)
			assert.Equal(t, tt.expect, o)
			assert.True(t, errors.Is(err, tt.expect_err))
		})
	}
}
