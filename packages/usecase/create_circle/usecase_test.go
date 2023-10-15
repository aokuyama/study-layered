package usecase_test

import (
	"errors"
	"testing"

	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	mock_owner "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/create_circle"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)
	cr.EXPECT().Save(gomock.Any()).Return(nil)

	u := New(or, cr)
	_, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})
	assert.NoError(t, err)
}

func TestOwnerIDInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_owner.NewMockOwnerRepository(ctrl)
	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateCircleInput{CircleName: "circle"})

	assert.Error(t, err)
	assert.Equal(t, "invalid UUID length: 0", err.Error())
	assert.Nil(t, out)
}

func TestOwnerNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, errors.New("owner not found"))
	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})

	assert.Error(t, err)
	assert.Equal(t, "owner not found", err.Error())
	assert.Nil(t, out)
}

func TestGenerateCircleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000"})

	assert.Error(t, err)
	assert.Equal(t, "can`t be blank", err.Error())
	assert.Nil(t, out)
}

func TestSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)
	cr.EXPECT().Save(gomock.Any()).Return(errors.New("save error"))

	u := New(or, cr)
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}
