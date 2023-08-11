package circle_test

import (
	"testing"

	. "aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"aokuyama/circle_scheduler-api/packages/domain/model/common"
	"aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	var e *Circle
	var err error
	i := util.PanicOr(common.GenerateUUID())
	e, err = NewCircle(*i)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.NoError(t, err)
}

func TestIdenticalEntity(t *testing.T) {
	i1 := util.PanicOr(common.GenerateUUID())
	e1 := util.PanicOr(NewCircle(*i1))
	e2 := util.PanicOr(NewCircle(*i1))
	i2 := util.PanicOr(common.GenerateUUID())
	e3 := util.PanicOr(NewCircle(*i2))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
