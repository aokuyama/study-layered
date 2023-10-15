package event_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

var circleID circle.CircleID

func TestEntity(t *testing.T) {
	var e *Event
	var err error
	n := "circle"
	e, err = GenerateEvent(&circleID, &n)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.Equal(t, "circle", e.Name.String())
	assert.NoError(t, err)
}

func TestIdenticalEntity(t *testing.T) {
	n := util.PanicOr(NewName("a"))
	p := util.PanicOr(path.GeneratePath())
	i1 := util.PanicOr(common.GenerateUUID())
	e1 := util.PanicOr(NewEvent(i1, &circleID, n, p))
	e2 := util.PanicOr(NewEvent(i1, &circleID, n, p))
	i2 := util.PanicOr(common.GenerateUUID())
	e3 := util.PanicOr(NewEvent(i2, &circleID, n, p))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
