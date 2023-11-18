package test

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

func GenGuest(n int) *guest.Guest {
	if n == 1 {
		return PanicOr(guest.NewGuest(P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), P("guest1"), P[uint8](1)))
	}
	if n == 11 {
		return PanicOr(guest.NewGuest(P("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"), P("guest1_1"), P[uint8](1)))
	}
	if n == 2 {
		return PanicOr(guest.NewGuest(P("d833a112-95e8-4042-ab02-ffde48bc874a"), P("guest2"), P[uint8](1)))
	}
	panic("undefined factory")
}
