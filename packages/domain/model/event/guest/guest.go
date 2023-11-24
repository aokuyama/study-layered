package guest

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type Guest struct {
	userId user.UserID
	name   string
	number uint8
}

type GuestInput struct {
	UserID string
	Name   string
	Number uint8
}

func NewGuest(i *GuestInput) (*Guest, error) {
	ID, err := user.NewUserID(i.UserID)
	if err != nil {
		return nil, err
	}
	if len(i.Name) <= 0 || len(i.Name) > 10 {
		return nil, errs.NewBadParam("must 1~10 characters")
	}
	if i.Number <= 0 || i.Number > 5 {
		return nil, errs.NewBadParam("be 1~5")
	}
	g := Guest{*ID, i.Name, i.Number}
	return &g, nil
}

func (g *Guest) UserID() *user.UserID {
	return &g.userId
}

func (g *Guest) Name() string {
	return g.name
}

func (g *Guest) Number() uint8 {
	return g.number
}

func (gn Guest) Identical(g *Guest) bool {
	return gn.UserID().Equals(g.UserID().UUID)
}

func (gn Guest) EqualsSafe(g *Guest) (bool, error) {
	if !gn.Identical(g) {
		return false, errs.NewBadParam("no identical entity")
	}
	if gn.name != g.name {
		return false, nil
	}
	if gn.number != g.number {
		return false, nil
	}
	return true, nil
}

func (gn Guest) Equals(g *Guest) bool {
	e, err := gn.EqualsSafe(g)
	if err != nil {
		panic(err)
	}
	return e
}
