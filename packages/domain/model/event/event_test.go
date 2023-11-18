package event_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/test"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	var e *Event
	var err error
	i := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	c := "d833a112-95e8-4042-ab02-ffde48bc874a"
	n := "event"
	p := path.Path{}
	e, err = NewEvent(&i, &c, &n, &p)
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", e.ID().String())
	assert.Equal(t, "d833a112-95e8-4042-ab02-ffde48bc874a", e.CircleID().String())
	assert.Equal(t, "event", e.Name().String())
	assert.NoError(t, err)
}

func TestErrorNewEntity(t *testing.T) {
	tests := []struct {
		testName           string
		id, circleID, name string
	}{
		{"id", "invalid", "d833a112-95e8-4042-ab02-ffde48bc874a", "circle"},
		{"circleID", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "invalid", "circle"},
		{"name", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "d833a112-95e8-4042-ab02-ffde48bc874a", ""},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			p := path.Path{}
			e, err := NewEvent(&tt.id, &tt.circleID, &tt.name, &p)
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
	p := path.Path{}

	e1 := util.PanicOr(NewEvent(&i1, &ci, &n, &p))
	e2 := util.PanicOr(NewEvent(&i1, &ci, &n, &p))
	e3 := util.PanicOr(NewEvent(&i2, &ci, &n, &p))
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
