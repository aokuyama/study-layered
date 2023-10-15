package event_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

var circleID circle.CircleID

func TestEntity(t *testing.T) {
	var e *EventEntity
	var err error
	n := util.PanicOr(NewName("event"))
	i := util.PanicOr(common.GenerateUUID())
	e, err = NewEventEntity(i, &circleID, n)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.Equal(t, "event", e.Name.String())
	assert.NoError(t, err)
}

func TestIdenticalEntity(t *testing.T) {
	n := util.PanicOr(NewName("a"))
	i1 := util.PanicOr(common.GenerateUUID())
	e1 := util.PanicOr(NewEventEntity(i1, &circleID, n))
	e2 := util.PanicOr(NewEventEntity(i1, &circleID, n))
	i2 := util.PanicOr(common.GenerateUUID())
	e3 := util.PanicOr(NewEventEntity(i2, &circleID, n))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
