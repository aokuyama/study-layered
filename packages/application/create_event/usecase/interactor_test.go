package usecase_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/create_event/usecase"
	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	mock_event "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/.mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_event.NewMockEventFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&event.EventEntity{}, nil)

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)
	cr.EXPECT().Create(gomock.Any()).Return(nil)

	u := New(f, or, cr)
	_, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})
	assert.NoError(t, err)
}

func TestCircleIDInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_event.NewMockEventFactory(ctrl)
	or := mock_circle.NewMockCircleRepository(ctrl)
	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateEventInput{EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "invalid UUID length: 0", err.Error())
	assert.Nil(t, out)
}

func TestCircleNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_event.NewMockEventFactory(ctrl)
	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, errors.New("circle not found"))
	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "circle not found", err.Error())
	assert.Nil(t, out)
}

func TestCreateEventError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_event.NewMockEventFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("create error"))

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000"})

	assert.Error(t, err)
	assert.Equal(t, "create error", err.Error())
	assert.Nil(t, out)
}

func TestCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := mock_event.NewMockEventFactory(ctrl)
	f.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&event.EventEntity{}, nil)

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)
	cr.EXPECT().Create(gomock.Any()).Return(errors.New("save error"))

	u := New(f, or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}
