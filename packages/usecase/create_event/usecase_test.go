package usecase_test

import (
	"errors"
	"testing"

	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	mock_event "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/create_event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)
	cr.EXPECT().Save(gomock.Any()).Return(nil)

	u := New(or, cr)
	_, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})
	assert.NoError(t, err)
}

func TestCircleIDInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_circle.NewMockCircleRepository(ctrl)
	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateEventInput{EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "invalid UUID length: 0", err.Error())
	assert.Nil(t, out)
}

func TestCircleNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, errors.New("circle not found"))
	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "circle not found", err.Error())
	assert.Nil(t, out)
}

func TestGenerateEventError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000"})

	assert.Error(t, err)
	assert.Equal(t, "can`t be blank", err.Error())
	assert.Nil(t, out)
}

func TestSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	or := mock_circle.NewMockCircleRepository(ctrl)
	or.EXPECT().Find(gomock.Any()).Return(nil, nil)

	cr := mock_event.NewMockEventRepository(ctrl)
	cr.EXPECT().Save(gomock.Any()).Return(errors.New("save error"))

	u := New(or, cr)
	out, err := u.Invoke(&CreateEventInput{CircleID: "550e8400-e29b-41d4-a716-446655440000", EventName: "event"})

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	assert.Nil(t, out)
}
