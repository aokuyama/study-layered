package test

import (
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

func GenGuest(n int) *Guest {
	if n == 1 {
		return PanicOr(NewGuest(&GuestInput{UserID: "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", Name: "guest1", Number: 1}))
	}
	if n == 11 {
		return PanicOr(NewGuest(&GuestInput{UserID: "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", Name: "guest1_1", Number: 1}))
	}
	if n == 2 {
		return PanicOr(NewGuest(&GuestInput{UserID: "d833a112-95e8-4042-ab02-ffde48bc874a", Name: "guest2", Number: 1}))
	}
	panic("undefined factory")
}
