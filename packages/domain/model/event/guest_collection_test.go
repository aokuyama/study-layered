package event_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestNewGuestCollection(t *testing.T) {
	c := NewEmptyGuestCollection()
	assert.True(t, c.Empty())
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
	assert.Equal(t, "guest1", *c1.Nth(0).Name())
	assert.Equal(t, "guest1_1", *c11.Nth(0).Name())
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
	assert.Equal(t, "guest1", *c1.Nth(0).Name())
	assert.Equal(t, "guest1_1", *c11.Nth(0).Name())
	c2 := c1.AppendOrUpdate(g2)
	assert.Equal(t, "guest2", *c2.Nth(1).Name())
}

func TestFailAppendOrUpdateGuest(t *testing.T) {
	g1 := test.GenGuest(1)

	c0 := NewEmptyGuestCollection()
	c1 := c0.AppendOrUpdate(g1)
	f1 := c1.AppendOrUpdate(g1)
	assert.Nil(t, f1, "not updated")
}
