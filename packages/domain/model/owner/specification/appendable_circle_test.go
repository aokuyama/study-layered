package specification_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/specification"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAppendableCircleSpec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		expect error
		mock   func(*mock_circle.MockCircleRepository)
	}{
		{"satisfied", nil, func(cr *mock_circle.MockCircleRepository) {
			cs := []circle.CircleID{}
			cr.EXPECT().SearchByOwner(gomock.Any()).Return(&cs, nil)
		}},
		{"load_error", errors.New("load_err"), func(cr *mock_circle.MockCircleRepository) {
			cs := []circle.CircleID{}
			cr.EXPECT().SearchByOwner(gomock.Any()).Return(&cs, errors.New("load_err"))
		}},
		{"circle full", errors.New("unable to append circle"), func(cr *mock_circle.MockCircleRepository) {
			cs := []circle.CircleID{{}, {}}
			cr.EXPECT().SearchByOwner(gomock.Any()).Return(&cs, nil)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := mock_circle.NewMockCircleRepository(ctrl)
			tt.mock(cr)
			u := NewAppendableCircleSpec(cr)
			err := u.IsSatisfiedBy(&owner.OwnerID{})
			assert.Equal(t, tt.expect, err)
		})
	}
}
