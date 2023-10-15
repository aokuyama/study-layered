package usecase_test

import (
	"errors"
	"testing"

	mock_owner "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/create_owner"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_owner.NewMockOwnerRepository(ctrl)
	r.EXPECT().Save(gomock.Any()).Return(nil)

	u := New(r)
	_, err := u.Invoke(&CreateOwnerInput{})

	assert.NoError(t, err)
}

func TestGenerateError(t *testing.T) {
	t.Skip()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_owner.NewMockOwnerRepository(ctrl)

	u := New(r)
	out, err := u.Invoke(&CreateOwnerInput{ /* error input */ })

	assert.Error(t, err)
	assert.Nil(t, out)
}

func TestSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_owner.NewMockOwnerRepository(ctrl)
	r.EXPECT().Save(gomock.Any()).Return(errors.New("save error"))

	u := New(r)
	out, err := u.Invoke(&CreateOwnerInput{})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}
