package user_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestIdenticalEntity(t *testing.T) {
	i1 := UserID{test.PanicOr(common.GenerateUUID())}
	e1 := test.PanicOr(NewUser(&i1))
	e2 := test.PanicOr(NewUser(&i1))
	i2 := UserID{test.PanicOr(common.GenerateUUID())}
	e3 := test.PanicOr(NewUser(&i2))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
