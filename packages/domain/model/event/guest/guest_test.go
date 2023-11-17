package guest_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestNewGuest(t *testing.T) {
	tests := []struct {
		testName string
		id, name string
		number   uint8
	}{
		{"1", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "guest", 1},
		{"2", "d833a112-95e8-4042-ab02-ffde48bc874a", "ok", 5},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			g, err := NewGuest(&tt.id, &tt.name, &tt.number)
			assert.NotNil(t, g)
			assert.NoError(t, err)
			assert.Equal(t, tt.id, g.UserID().String())
			assert.Equal(t, tt.name, *g.Name())
			assert.Equal(t, tt.number, *g.Number())
		})
	}
}

func TestErrorNewGuest(t *testing.T) {
	tests := []struct {
		testName string
		id, name string
		number   uint8
	}{
		{"id", "invalid", "g", 1},
		{"name", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "", 1},
		{"number", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "g", 0},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			g, err := NewGuest(&tt.id, &tt.name, &tt.number)
			assert.Nil(t, g)
			assert.Error(t, err)
			assert.True(t, errors.Is(err, errs.ErrBadParam))
		})
	}
}

func TestIdentical(t *testing.T) {
	sample := util.PanicOr(NewGuest(util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), util.P("g"), util.P[uint8](1)))
	g1 := util.PanicOr(NewGuest(util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), util.P("g"), util.P[uint8](1)))
	g2 := util.PanicOr(NewGuest(util.P("d833a112-95e8-4042-ab02-ffde48bc874a"), util.P("g"), util.P[uint8](1)))
	assert.True(t, sample.Identical(g1))
	assert.False(t, sample.Identical(g2))
}

func TestEquals(t *testing.T) {
	sample := util.PanicOr(NewGuest(util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), util.P("g"), util.P[uint8](1)))
	tests := []struct {
		testName string
		id, name string
		number   uint8
		equals   bool
		err      error
	}{
		{"equals", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "g", 1, true, nil},
		{"name not equals", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "gg", 1, false, nil},
		{"number not equals", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "g", 2, false, nil},
		{"id not equals error", "d833a112-95e8-4042-ab02-ffde48bc874a", "g", 1, false, errs.ErrBadParam},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			g := util.PanicOr(NewGuest(&tt.id, &tt.name, &tt.number))
			e, err := sample.Equals(g)
			assert.Equal(t, tt.equals, e)
			assert.True(t, errors.Is(err, tt.err))
		})
	}
}
