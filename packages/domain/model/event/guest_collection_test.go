package event_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"

	"github.com/stretchr/testify/assert"
)

func TestNewGuestCollection(t *testing.T) {
	c := NewEmptyGuestCollection()
	assert.True(t, c.IsEmpty())
}
