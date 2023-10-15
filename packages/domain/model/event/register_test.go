package event_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	var circleID circle.CircleID
	var e *RegisterEvent
	var err error
	n := "circle"
	e, err = GenerateRegisterEvent(&circleID, &n)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.Equal(t, "circle", e.Name.String())
	assert.NoError(t, err)
}
