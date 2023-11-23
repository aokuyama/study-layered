package usecase_test

import (
	"errors"
	"fmt"
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

	i := UserJoinToEventInput{e.ID().String(), g.UserID().String(), *g.Name(), *g.Number()}
	ei := test.PanicOr(event.NewEventID(e.ID().String()))

	equalEvent := e.JoinGuest(g)
	diffEvent := test.GenEvent(11)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		mock_repository func(r *mock_event.MockEventRepository)
		input           UserJoinToEventInput
	}{
		{"success update", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(diffEvent, nil)
			r.EXPECT().Update(gomock.Any(), diffEvent).Return(nil)
		}, i},

		{"success no update", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(equalEvent, nil)
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
		})
	}
}

func TestInvokeError(t *testing.T) {
	diffEvent := test.GenEvent(11)

	i := UserJoinToEventInput{"26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "d833a112-95e8-4042-ab02-ffde48bc874a", "name", 1}
	ei := test.PanicOr(event.NewEventID("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"))

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
			r.EXPECT().Find(ei).Return(nil, fmt.Errorf("test: %w", errs.ErrNotFound))
		}, i, errs.ErrNotFound},

		{"fail update event", func(r *mock_event.MockEventRepository) {
			r.EXPECT().Find(ei).Return(diffEvent, nil)
			r.EXPECT().Update(gomock.Any(), diffEvent).Return(fmt.Errorf("test: %w", errs.ErrFatal))
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
