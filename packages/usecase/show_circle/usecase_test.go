package usecase_test

import (
	"errors"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	mock_circle "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle/.mock"
	. "github.com/aokuyama/circle_scheduler-api/packages/usecase/show_circle"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInvoke(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_circle.NewMockCircleRepository(ctrl)
	c := &circle.CircleEntity{}
	cr.EXPECT().FindByPath(gomock.Any()).Return(c, nil)

	u := New(cr)
	out, err := u.Invoke(&ShowCircleInput{Path: "0123456789abcdef"})
	assert.NoError(t, err)
	assert.Equal(t, c, out.Circle)
}

func TestPathInputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_circle.NewMockCircleRepository(ctrl)

	u := New(cr)
	out, err := u.Invoke(&ShowCircleInput{Path: "invalid"})

	assert.Equal(t, "must 16 characters", err.Error())
	assert.Nil(t, out)
}

func TestCircleNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cr := mock_circle.NewMockCircleRepository(ctrl)
	cr.EXPECT().FindByPath(gomock.Any()).Return(nil, errors.New("not found circle"))

	u := New(cr)
	out, err := u.Invoke(&ShowCircleInput{Path: "0123456789abcdef"})

	assert.Equal(t, "not found circle", err.Error())
	assert.Nil(t, out)
}
