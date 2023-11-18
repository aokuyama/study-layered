package test

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"
)

func GenGuest(n int) *guest.Guest {
	if n == 1 {
		return util.PanicOr(guest.NewGuest(util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), util.P("guest1"), util.P[uint8](1)))
	}
	if n == 11 {
		return util.PanicOr(guest.NewGuest(util.P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), util.P("guest1_1"), util.P[uint8](1)))
	}
	if n == 2 {
		return util.PanicOr(guest.NewGuest(util.P("d833a112-95e8-4042-ab02-ffde48bc874a"), util.P("guest2"), util.P[uint8](1)))
	}
	panic("undefined factory")
}
