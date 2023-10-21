package usecase_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	mock_owner "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/create_owner"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_owner.NewMockOwnerFactory(ctrl)
	f.EXPECT().Create().Return(&owner.Owner{}, nil)

	r := mock_owner.NewMockOwnerRepository(ctrl)
	r.EXPECT().Save(gomock.Any()).Return(nil)

	u := New(f, r)
	_, err := u.Invoke(&CreateOwnerInput{})

	assert.NoError(t, err)
}

func TestFactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_owner.NewMockOwnerFactory(ctrl)
	f.EXPECT().Create().Return(nil, errors.New("create error"))

	r := mock_owner.NewMockOwnerRepository(ctrl)

	u := New(f, r)
	out, err := u.Invoke(&CreateOwnerInput{})

	assert.Error(t, err)
	assert.Equal(t, "create error", err.Error())
	assert.Nil(t, out)
}

func TestSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_owner.NewMockOwnerFactory(ctrl)
	f.EXPECT().Create().Return(&owner.Owner{}, nil)

	r := mock_owner.NewMockOwnerRepository(ctrl)
	r.EXPECT().Save(gomock.Any()).Return(errors.New("save error"))

	u := New(f, r)
	out, err := u.Invoke(&CreateOwnerInput{})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}
