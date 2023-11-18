package owner_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestIdenticalEntity(t *testing.T) {
	i1 := OwnerID{test.PanicOr(common.GenerateUUID())}
	e1 := test.PanicOr(NewOwner(&i1))
	e2 := test.PanicOr(NewOwner(&i1))
	i2 := OwnerID{test.PanicOr(common.GenerateUUID())}
	e3 := test.PanicOr(NewOwner(&i2))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
