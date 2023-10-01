package owner_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	var e *Owner
	var err error
	e, err = GenerateOwner()
	assert.Equal(t, 36, len(e.ID.String()))
	assert.NoError(t, err)
}

func TestIdenticalEntity(t *testing.T) {
	i1 := util.PanicOr(common.GenerateUUID())
	e1 := util.PanicOr(NewOwner(i1))
	e2 := util.PanicOr(NewOwner(i1))
	i2 := util.PanicOr(common.GenerateUUID())
	e3 := util.PanicOr(NewOwner(i2))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
