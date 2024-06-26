package event_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	var e *Event
	var err error
	i := test.GenUUIDStinrg(1)
	c := test.GenUUIDStinrg(2)
	n := "event"
	p := test.GenPathString()
	g := []guest.GuestInput{{
		UserID: test.GenUUIDStinrg(3), Name: "g1", Number: 1,
	}}

	e, err = NewEvent(&EventInput{i, c, n, p, g})
	assert.Equal(t, test.GenUUIDStinrg(1), e.ID().String())
	assert.Equal(t, test.GenUUIDStinrg(2), e.CircleID().String())
	assert.Equal(t, "event", e.Name().String())
	assert.Equal(t, test.GenUUIDStinrg(3), e.Guest().Nth(0).UserID().String())
	assert.NoError(t, err)
}

func TestErrorNewEntity(t *testing.T) {
	tests := []struct {
		testName           string
		id, circleID, name string
		guest              []guest.GuestInput
	}{
		{"id", "invalid", "d833a112-95e8-4042-ab02-ffde48bc874a", "circle", []guest.GuestInput{}},
		{"circleID", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "invalid", "circle", []guest.GuestInput{}},
		{"name", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "d833a112-95e8-4042-ab02-ffde48bc874a", "", []guest.GuestInput{}},
		{"guest", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "d833a112-95e8-4042-ab02-ffde48bc874a", "circle", []guest.GuestInput{{UserID: "", Name: "g1", Number: 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			e, err := NewEvent(&EventInput{tt.id, tt.circleID, tt.name, test.GenPathString(), tt.guest})
			assert.Nil(t, e)
			assert.Error(t, err)
		})
	}
}

func TestIdenticalEntity(t *testing.T) {
	n := "a"
	i1 := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	i2 := "550e8400-e29b-41d4-a716-446655440000"
	ci := "d833a112-95e8-4042-ab02-ffde48bc874a"
	p := test.GenPathString()
	g := []guest.GuestInput{}

	e1 := test.PanicOr(NewEvent(&EventInput{i1, ci, n, p, g}))
	e2 := test.PanicOr(NewEvent(&EventInput{i1, ci, n, p, g}))
	e3 := test.PanicOr(NewEvent(&EventInput{i2, ci, n, p, g}))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}

func TestJoinGuest(t *testing.T) {
	e := test.GenEvent(1)
	g1 := test.GenGuest(1)
	g11 := test.GenGuest(11)
	e1 := e.JoinGuest(g1)
	assert.Equal(t, 0, e.Guest().Len())
	assert.Equal(t, 1, e1.Guest().Len())
	e11 := e1.JoinGuest(g11)
	assert.Equal(t, 1, e11.Guest().Len())
	f1 := e1.JoinGuest(g1)
	assert.Nil(t, f1)
}

func TestRemoveGuest(t *testing.T) {
	g1 := test.GenGuest(1)
	g2 := test.GenGuest(2)
	e12 := test.GenEvent(1).JoinGuest(g1).JoinGuest(g2)
	assert.Equal(t, 2, e12.Guest().Len())
	e2 := e12.RemoveGuest(g1.UserID())
	assert.Equal(t, 1, e2.Guest().Len())
	assert.Nil(t, e2.RemoveGuest(g1.UserID()))
	e0 := e2.RemoveGuest(g2.UserID())
	assert.Equal(t, 0, e0.Guest().Len())
}
