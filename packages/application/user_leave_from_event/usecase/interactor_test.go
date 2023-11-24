package usecase_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/user_leave_from_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	mock_event "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/.mock"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	e := test.GenEvent(1)
	g1 := test.GenGuest(1)

	i := UserLeaveFromEventInput{e.ID().String(), g1.UserID().String()}
	ei := test.PanicOr(event.NewEventID(e.ID().String()))

	before := test.GenEvent(1).JoinGuest(g1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_event.MockEventRepository)
		input           UserLeaveFromEventInput
	}{
		{"success update", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(before, nil)
			r.EXPECT().Update(gomock.Any(), before).Return(nil)
		}, i},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mock_event.NewMockEventRepository(ctrl)
			tt.mock_repository(r)

			u := New(r)
			o, err := u.Invoke(&tt.input)
			assert.NotNil(t, o)
			assert.NoError(t, err)
			assert.NotEqual(t, &o.Event, before)
		})
	}
}

func TestInvokeError(t *testing.T) {
	e := test.GenEvent(1)
	g := test.GenGuest(1)

	i := UserLeaveFromEventInput{e.ID().String(), g.UserID().String()}
	ei := test.PanicOr(event.NewEventID(e.ID().String()))

	event := e.JoinGuest(g)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_event.MockEventRepository)
		input           UserLeaveFromEventInput
		expect          error
	}{
		{"invalid guest id", func(r *mock_event.MockEventRepository) {
		}, UserLeaveFromEventInput{"26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "fail"}, errs.ErrBadParam},

		{"invalid event id", func(r *mock_event.MockEventRepository) {
		}, UserLeaveFromEventInput{"fail", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"}, errs.ErrBadParam},

		{"event not found", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(nil, errs.NewNotFound("test"))
		}, i, errs.ErrNotFound},

		{"guest not found", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(e, nil)
		}, i, errs.ErrNotFound},

		{"fail update event", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(event, nil)
			r.EXPECT().Update(gomock.Any(), event).Return(errs.NewFatal("test"))
		}, i, errs.ErrFatal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mock_event.NewMockEventRepository(ctrl)
			tt.mock_repository(r)

			u := New(r)
			o, err := u.Invoke(&tt.input)
			assert.Nil(t, o)
			assert.True(t, errors.Is(err, tt.expect))
		})
	}
}
