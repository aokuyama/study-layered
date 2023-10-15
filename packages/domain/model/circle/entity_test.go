package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

var ownerID owner.OwnerID

func TestEntity(t *testing.T) {
	var e *CircleEntity
	var err error
	i := util.PanicOr(common.GenerateUUID())
	n := util.PanicOr(NewName("circle"))

	e, err = NewCircleEntity(i, &ownerID, n)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.Equal(t, "circle", e.Name.String())
	assert.NoError(t, err)
}

func TestIdenticalEntity(t *testing.T) {
	n := util.PanicOr(NewName("a"))
	i1 := util.PanicOr(common.GenerateUUID())
	e1 := util.PanicOr(NewCircleEntity(i1, &ownerID, n))
	e2 := util.PanicOr(NewCircleEntity(i1, &ownerID, n))
	i2 := util.PanicOr(common.GenerateUUID())
	e3 := util.PanicOr(NewCircleEntity(i2, &ownerID, n))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
