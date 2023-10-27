package event_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	var e *EventEntity
	var err error
	i := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	c := "d833a112-95e8-4042-ab02-ffde48bc874a"
	n := "event"
	e, err = NewEventEntity(&i, &c, &n)
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
			e, err := NewEventEntity(&tt.id, &tt.circleID, &tt.name)
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

	e1 := util.PanicOr(NewEventEntity(&i1, &ci, &n))
	e2 := util.PanicOr(NewEventEntity(&i1, &ci, &n))
	e3 := util.PanicOr(NewEventEntity(&i2, &ci, &n))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
