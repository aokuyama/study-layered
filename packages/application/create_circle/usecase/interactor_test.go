package usecase

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	mock_owner "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/.mock"
	mock_specification "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/specification/.mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&circle.CircleEntity{}, nil)

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)
	cr.EXPECT().Create(gomock.Any()).Return(nil)

	s := mock_specification.NewMockAppendableCircleSpec(ctrl)
	s.EXPECT().IsSatisfiedBy(gomock.Any()).Return(nil)

	u := createCircleInteractor{f, or, cr, s}
	_, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})
	assert.NoError(t, err)
}

func TestOwnerIDInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)
	or := mock_owner.NewMockOwnerRepository(ctrl)
	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateCircleInput{CircleName: "circle"})

	assert.Error(t, err)
	assert.True(t, errors.Is(err, errs.ErrBadParam))
	assert.Nil(t, out)
}

func TestOwnerNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)
	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, errors.New("owner not found"))
	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})

	assert.Error(t, err)
	assert.Equal(t, "owner not found", err.Error())
	assert.Nil(t, out)
}

func TestCreateCircleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("create error"))

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)

	s := mock_specification.NewMockAppendableCircleSpec(ctrl)
	s.EXPECT().IsSatisfiedBy(gomock.Any()).Return(nil)

	u := createCircleInteractor{f, or, cr, s}
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000"})

	assert.Error(t, err)
	assert.Equal(t, "create error", err.Error())
	assert.Nil(t, out)
}

func TestCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&circle.CircleEntity{}, nil)

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)
	cr.EXPECT().Create(gomock.Any()).Return(errors.New("save error"))

	s := mock_specification.NewMockAppendableCircleSpec(ctrl)
	s.EXPECT().IsSatisfiedBy(gomock.Any()).Return(nil)

	u := createCircleInteractor{f, or, cr, s}
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}

func TestUnableAppendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_circle.NewMockCircleFactory(ctrl)

	or := mock_owner.NewMockOwnerRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_circle.NewMockCircleRepository(ctrl)

	s := mock_specification.NewMockAppendableCircleSpec(ctrl)
	s.EXPECT().IsSatisfiedBy(gomock.Any()).Return(errors.New("unable to append circle"))

	u := createCircleInteractor{f, or, cr, s}
	out, err := u.Invoke(&CreateCircleInput{OwnerID: "550e8400-e29b-41d4-a716-446655440000", CircleName: "circle"})
	assert.Error(t, err)
	assert.Equal(t, "unable to append circle", err.Error())
	assert.Nil(t, out)
}
