package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	var ownerID owner.OwnerID
	n := "circle"
	e, err := GenerateRegisterCircle(&ownerID, &n)
	assert.Equal(t, 36, len(e.ID.String()))
	assert.Equal(t, "circle", e.Name.String())
	assert.NoError(t, err)
}
