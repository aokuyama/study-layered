package guest

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type Guest struct {
	userId user.UserID
	name   string
	number uint8
}

func NewGuest(userID *string, name *string, number *uint8) (*Guest, error) {
	i, err := user.NewUserID(*userID)
	if err != nil {
		return nil, err
	}
	if len(*name) <= 0 || len(*name) > 10 {
		return nil, fmt.Errorf("%w must 1~10 characters", errs.ErrBadParam)
	}
	if *number <= 0 || *number > 5 {
		return nil, fmt.Errorf("%w be 1~5", errs.ErrBadParam)
	}
	g := Guest{*i, *name, *number}
	return &g, nil
}

func (g *Guest) UserID() *user.UserID {
	return &g.userId
}

func (g *Guest) Name() *string {
	return &g.name
}

func (g *Guest) Number() *uint8 {
	return &g.number
}
