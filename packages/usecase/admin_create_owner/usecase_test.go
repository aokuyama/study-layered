package admin_create_owner_test

import (
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	mock_owner "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/admin_create_owner"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_owner.NewMockOwnerRepository(ctrl)
	o := owner.Owner{}
	r.EXPECT().Save(gomock.Any()).Return(&o, nil)

	u := New(r)
	out, err := u.Invoke(&Input{})

	assert.NoError(t, err)
	assert.Equal(t, &o, out.Owner)
}
