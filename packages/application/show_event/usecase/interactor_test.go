package usecase_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/show_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	mock_event "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/.mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_event.NewMockEventRepository(ctrl)
	c := &event.EventEntity{}
	cr.EXPECT().FindByPath(gomock.Any()).Return(c, nil)

	u := New(cr)
	out, err := u.Invoke(&ShowEventInput{Path: "0123456789abcdef"})
	assert.NoError(t, err)
	assert.Equal(t, c, out.Event)
}

func TestPathInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_event.NewMockEventRepository(ctrl)

	u := New(cr)
	out, err := u.Invoke(&ShowEventInput{Path: "invalid"})

	assert.Error(t, err)
	assert.True(t, errors.Is(err, errs.ErrBadParam))
	assert.Nil(t, out)
}

func TestEventNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_event.NewMockEventRepository(ctrl)
	cr.EXPECT().FindByPath(gomock.Any()).Return(nil, errors.New("not found event"))

	u := New(cr)
	out, err := u.Invoke(&ShowEventInput{Path: "0123456789abcdef"})

	assert.Equal(t, "not found event", err.Error())
	assert.Nil(t, out)
}
