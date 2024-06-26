package usecase_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/application/user_join_to_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	mock_event "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/.mock"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	e := test.GenEvent(1)
	g := test.GenGuest(1)

	i := UserJoinToEventInput{e.ID().String(), g.UserID().String(), g.Name(), g.Number()}
	ei := test.PanicOr(event.NewEventID(e.ID().String()))

	before := e

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_event.MockEventRepository)
		input           UserJoinToEventInput
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

	i := UserJoinToEventInput{e.ID().String(), g.UserID().String(), g.Name(), g.Number()}
	ei := test.PanicOr(event.NewEventID(e.ID().String()))

	diffEvent := test.GenEvent(11)
	equalEvent := e.JoinGuest(g)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_event.MockEventRepository)
		input           UserJoinToEventInput
		expect          error
	}{
		{"invalid guest", func(r *mock_event.MockEventRepository) {
		}, UserJoinToEventInput{"26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "fail", "name", 1}, errs.ErrBadParam},

		{"invalid event id", func(r *mock_event.MockEventRepository) {
		}, UserJoinToEventInput{"fail", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "name", 1}, errs.ErrBadParam},

		{"event not found", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(nil, errs.NewNotFound("test"))
		}, i, errs.ErrNotFound},

		{"already appended", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(equalEvent, nil)
		}, i, errs.ErrConflict},

		{"fail update event", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(diffEvent, nil)
			r.EXPECT().Update(gomock.Any(), diffEvent).Return(errs.NewFatal("test"))
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
