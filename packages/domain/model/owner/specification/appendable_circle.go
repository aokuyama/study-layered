//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package specification

import (
	"errors"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner/collection"
)

type AppendableCircleSpec interface {
	IsSatisfiedBy(i *owner.OwnerID) error
}

type appendableCircleSpec struct {
	circleRepository circle.CircleRepository
}

func NewAppendableCircleSpec(cr circle.CircleRepository) *appendableCircleSpec {
	s := appendableCircleSpec{cr}
	return &s
}

func (s *appendableCircleSpec) IsSatisfiedBy(i *owner.OwnerID) error {
	cl, err := s.circleRepository.SearchByOwner(i)
	if err != nil {
		return err
	}
	cc := collection.NewOwnerCircles(*cl)
	if !cc.IsAppendable() {
		return errors.New("unable to append circle")
	}
	return nil
}
