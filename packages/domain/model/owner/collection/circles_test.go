package collection_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/collection"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"
	"github.com/stretchr/testify/assert"
)

func TestOwnerCirclesIsAppendable(t *testing.T) {
	u := circle.CircleID{UUID: test.PanicOr(common.GenerateUUID())}
	tests := []struct {
		name   string
		expect bool
		list   []circle.CircleID
	}{
		{"0", true, []circle.CircleID{}},
		{"1", false, []circle.CircleID{u}},
		{">1", false, []circle.CircleID{u, u}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewOwnerCircles(tt.list)
			assert.Equal(t, tt.expect, c.IsAppendable())
		})
	}
}
