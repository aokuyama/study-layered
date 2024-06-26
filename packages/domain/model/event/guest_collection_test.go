package event_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestNewGuestCollection(t *testing.T) {
	tests := []struct {
		name  string
		len   int
		input []guest.GuestInput
	}{
		{"empty", 0, []guest.GuestInput{}},
		{"one", 1, []guest.GuestInput{
			{UserID: test.GenUUIDStinrg(1), Name: "a", Number: 1},
		}},
		{"many", 3, []guest.GuestInput{
			{UserID: test.GenUUIDStinrg(1), Name: "a", Number: 1},
			{UserID: test.GenUUIDStinrg(2), Name: "a", Number: 1},
			{UserID: test.GenUUIDStinrg(3), Name: "a", Number: 1},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewGuestCollection(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.len, c.Len())
		})
	}
}

func TestNewEmptyGuestCollection(t *testing.T) {
	c := NewEmptyGuestCollection()
	assert.Equal(t, 0, c.Len())
	assert.True(t, c.Empty())
}

func TestErrorNewGuestCollection(t *testing.T) {
	tests := []struct {
		name  string
		input []guest.GuestInput
		err   error
	}{
		{"bad param", []guest.GuestInput{
			{UserID: "", Name: "", Number: 1},
		}, errs.ErrBadParam},
		{"dupuricate", []guest.GuestInput{
			{UserID: test.GenUUIDStinrg(1), Name: "a", Number: 1},
			{UserID: test.GenUUIDStinrg(1), Name: "aa", Number: 1},
		}, errs.ErrFatal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewGuestCollection(tt.input)
			assert.Error(t, err)
			assert.Nil(t, c)
			assert.True(t, errors.Is(err, tt.err))
		})
	}
}

func TestAppendGuest(t *testing.T) {
	g1 := test.GenGuest(1)
	g2 := test.GenGuest(2)

	c0 := NewEmptyGuestCollection()
	c1 := c0.Append(g1)
	assert.Equal(t, 0, c0.Len())
	assert.Equal(t, 1, c1.Len())
	c2 := c1.Append(g2)
	assert.Equal(t, 0, c0.Len())
	assert.Equal(t, 1, c1.Len())
	assert.Equal(t, 2, c2.Len())
	assert.True(t, g1.Identical(c1.Nth(0)))
	assert.True(t, g1.Identical(c2.Nth(0)))
	assert.True(t, g2.Identical(c2.Nth(1)))
}

func TestFailAppendGuest(t *testing.T) {
	g1 := test.GenGuest(1)
	g2 := test.GenGuest(2)

	c0 := NewEmptyGuestCollection()
	c1 := c0.Append(g1)
	c2 := c1.Append(g2)
	f := c1.Append(g1)
	assert.Nil(t, f)

	f2 := c2.Append(g1)
	f3 := c2.Append(g2)
	assert.Nil(t, f2)
	assert.Nil(t, f3)
	assert.Equal(t, 1, c1.Len())
	assert.Equal(t, 2, c2.Len())
}

func TestUpdateGuest(t *testing.T) {
	g1 := test.GenGuest(1)
	g11 := test.GenGuest(11)

	c0 := NewEmptyGuestCollection()
	c1 := c0.Append(g1)
	c11 := c1.Update(g11)
	assert.Equal(t, 1, c1.Len())
	assert.Equal(t, 1, c11.Len())
	assert.Equal(t, "guest1", c1.Nth(0).Name())
	assert.Equal(t, "guest1_1", c11.Nth(0).Name())
}

func TestFailUpdate(t *testing.T) {
	g1 := test.GenGuest(1)
	g2 := test.GenGuest(2)

	c0 := NewEmptyGuestCollection()
	c1 := c0.Append(g1)

	f1 := c1.Update(g2)
	assert.Nil(t, f1, "not found identical")

	f2 := c1.Update(g1)
	assert.Nil(t, f2, "not updated")
}

func TestAppendOrUpdateGuest(t *testing.T) {
	g1 := test.GenGuest(1)
	g11 := test.GenGuest(11)
	g2 := test.GenGuest(2)

	c0 := NewEmptyGuestCollection()
	c1 := c0.AppendOrUpdate(g1)
	c11 := c1.AppendOrUpdate(g11)
	assert.Equal(t, 1, c1.Len())
	assert.Equal(t, 1, c11.Len())
	assert.Equal(t, "guest1", c1.Nth(0).Name())
	assert.Equal(t, "guest1_1", c11.Nth(0).Name())
	c2 := c1.AppendOrUpdate(g2)
	assert.Equal(t, "guest2", c2.Nth(1).Name())
}

func TestFailAppendOrUpdateGuest(t *testing.T) {
	g1 := test.GenGuest(1)

	c0 := NewEmptyGuestCollection()
	c1 := c0.AppendOrUpdate(g1)
	f1 := c1.AppendOrUpdate(g1)
	assert.Nil(t, f1, "not updated")
}

func TestIdenticalItem(t *testing.T) {
	g1 := test.GenGuest(1)
	g11 := test.GenGuest(11)
	g2 := test.GenGuest(2)

	c0 := NewEmptyGuestCollection()
	c1 := c0.AppendOrUpdate(g1)
	c2 := c1.AppendOrUpdate(g11)
	assert.Equal(t, g1, c1.IdenticalItem(g1))
	assert.True(t, c1.ExistsIdentical(g1))

	assert.Equal(t, g11, c2.IdenticalItem(g1))
	assert.True(t, c2.ExistsIdentical(g1))

	assert.Nil(t, c1.IdenticalItem(g2))
	assert.False(t, c1.ExistsIdentical(g2))
}

func TestRemove(t *testing.T) {
	g1 := test.GenGuest(1)
	g2 := test.GenGuest(2)
	gc12 := NewEmptyGuestCollection().Append(g1).Append(g2)
	assert.Equal(t, 2, gc12.Len())

	gc2 := gc12.Remove(g1.UserID())
	gc1 := gc12.Remove(g2.UserID())

	for _, tt := range []struct {
		name                  string
		guest                 *GuestCollection
		existsID, notExistsID *user.UserID
	}{
		{"1", gc1, g1.UserID(), g2.UserID()},
		{"2", gc2, g2.UserID(), g1.UserID()},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, 1, tt.guest.Len())
			assert.Equal(t, tt.existsID.String(), tt.guest.Nth(0).UserID().String())
			assert.Nil(t, tt.guest.Remove(tt.notExistsID))
			assert.Equal(t, 0, tt.guest.Remove(tt.existsID).Len())
		})
	}
}
