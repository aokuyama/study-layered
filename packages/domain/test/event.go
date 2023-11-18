package test

import (
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
)

var event1Path = GenPathString()

func GenEvent(n int) *Event {
	if n == 1 {
		return PanicOr(NewEvent(&EventInput{
			ID:       "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1",
			CircleID: "d833a112-95e8-4042-ab02-ffde48bc874a",
			Name:     "event1",
			Path:     event1Path,
		}))
	}
	if n == 11 {
		return PanicOr(NewEvent(&EventInput{
			ID:       "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1",
			CircleID: "d833a112-95e8-4042-ab02-ffde48bc874a",
			Name:     "event1_1",
			Path:     event1Path,
		}))
	}
	panic("undefined factory")
}
