package event_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	var circleID circle.CircleID
	n := "event"
	e, err := EventFactoryImpl{}.Create(&circleID, &n)
	assert.Equal(t, 36, len(e.ID().String()))
	assert.Equal(t, "event", e.Name().String())
	assert.True(t, e.Guest().IsEmpty())
	assert.NoError(t, err)
}
